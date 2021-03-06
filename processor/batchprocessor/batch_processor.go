// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package batchprocessor

import (
	"context"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"go.uber.org/zap"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.opentelemetry.io/collector/consumer/pdatautil"
	"go.opentelemetry.io/collector/internal/data"
	"go.opentelemetry.io/collector/processor"
)

// batch_processor is a component that accepts spans and metrics, places them
// into batches and sends downstream.
//
// batch_processor implements consumer.TraceConsumer and consumer.MetricsConsumer
//
// Batches are sent out with any of the following conditions:
// - batch size reaches cfg.SendBatchSize
// - cfg.Timeout is elapsed since the timestamp when the previous batch was sent out.
type batchProcessor struct {
	name   string
	logger *zap.Logger

	sendBatchSize uint32
	timeout       time.Duration

	timer *time.Timer
	done  chan struct{}

	newItem chan interface{}
	batch   batch
}

type batch interface {
	// export the current batch
	export(ctx context.Context) error

	// itemCount returns the size of the current batch
	itemCount() uint32

	// reset the current batch structure with zero/empty values.
	reset()

	// add item to the current batch
	add(item interface{})
}

var _ consumer.TraceConsumer = (*batchProcessor)(nil)
var _ consumer.MetricsConsumer = (*batchProcessor)(nil)

func newBatchProcessor(params component.ProcessorCreateParams, cfg *Config, batch batch) *batchProcessor {
	return &batchProcessor{
		name:   cfg.Name(),
		logger: params.Logger,

		sendBatchSize: cfg.SendBatchSize,
		timeout:       cfg.Timeout,
		done:          make(chan struct{}),
		newItem:       make(chan interface{}, 1),
		batch:         batch,
	}
}

func (bp *batchProcessor) GetCapabilities() component.ProcessorCapabilities {
	return component.ProcessorCapabilities{MutatesConsumedData: true}
}

// Start is invoked during service startup.
func (bp *batchProcessor) Start(context.Context, component.Host) error {
	go bp.startProcessingCycle()
	return nil
}

// Shutdown is invoked during service shutdown.
func (bp *batchProcessor) Shutdown(context.Context) error {
	close(bp.done)
	return nil
}

func (bp *batchProcessor) startProcessingCycle() {
	bp.timer = time.NewTimer(bp.timeout)
	for {
		select {
		case item := <-bp.newItem:
			bp.batch.add(item)

			if bp.batch.itemCount() >= bp.sendBatchSize {
				bp.timer.Stop()
				bp.sendItems(statBatchSizeTriggerSend)
				bp.resetTimer()
			}
		case <-bp.timer.C:
			if bp.batch.itemCount() > 0 {
				bp.sendItems(statTimeoutTriggerSend)
			}
			bp.resetTimer()
		case <-bp.done:
			if bp.batch.itemCount() > 0 {
				// TODO: Set a timeout on sendTraces or
				// make it cancellable using the context that Shutdown gets as a parameter
				bp.sendItems(statTimeoutTriggerSend)
			}
			return
		}
	}
}

func (bp *batchProcessor) resetTimer() {
	bp.timer.Reset(bp.timeout)
}

func (bp *batchProcessor) sendItems(measure *stats.Int64Measure) {
	// Add that it came form the trace pipeline?
	statsTags := []tag.Mutator{tag.Insert(processor.TagProcessorNameKey, bp.name)}
	_ = stats.RecordWithTags(context.Background(), statsTags, measure.M(1), statBatchSendSize.M(int64(bp.batch.itemCount())))

	if err := bp.batch.export(context.Background()); err != nil {
		bp.logger.Warn("Sender failed", zap.Error(err))
	}
	bp.batch.reset()
}

// ConsumeTraces implements TraceProcessor
func (bp *batchProcessor) ConsumeTraces(_ context.Context, td pdata.Traces) error {
	bp.newItem <- td
	return nil
}

// ConsumeTraces implements MetricsProcessor
func (bp *batchProcessor) ConsumeMetrics(_ context.Context, md pdata.Metrics) error {
	// First thing is convert into a different internal format
	bp.newItem <- md
	return nil
}

// newBatchTracesProcessor creates a new batch processor that batches traces by size or with timeout
func newBatchTracesProcessor(params component.ProcessorCreateParams, trace consumer.TraceConsumer, cfg *Config) *batchProcessor {
	return newBatchProcessor(params, cfg, newBatchTraces(trace))
}

// newBatchMetricsProcessor creates a new batch processor that batches metrics by size or with timeout
func newBatchMetricsProcessor(params component.ProcessorCreateParams, metrics consumer.MetricsConsumer, cfg *Config) *batchProcessor {
	return newBatchProcessor(params, cfg, newBatchMetrics(metrics))
}

type batchTraces struct {
	nextConsumer consumer.TraceConsumer
	traceData    pdata.Traces
	spanCount    uint32
}

func newBatchTraces(nextConsumer consumer.TraceConsumer) *batchTraces {
	b := &batchTraces{nextConsumer: nextConsumer}
	b.reset()
	return b
}

// add updates current batchTraces by adding new TraceData object
func (bt *batchTraces) add(item interface{}) {
	td := item.(pdata.Traces)
	newSpanCount := td.SpanCount()
	if newSpanCount == 0 {
		return
	}

	bt.spanCount += uint32(newSpanCount)
	td.ResourceSpans().MoveAndAppendTo(bt.traceData.ResourceSpans())
}

func (bt *batchTraces) export(ctx context.Context) error {
	return bt.nextConsumer.ConsumeTraces(ctx, bt.traceData)
}

func (bt *batchTraces) itemCount() uint32 {
	return bt.spanCount
}

// resets the current batchTraces structure with zero values
func (bt *batchTraces) reset() {
	bt.traceData = pdata.NewTraces()
	bt.spanCount = 0
}

type batchMetrics struct {
	nextConsumer consumer.MetricsConsumer
	metricData   data.MetricData
	metricCount  uint32
}

func newBatchMetrics(nextConsumer consumer.MetricsConsumer) *batchMetrics {
	b := &batchMetrics{nextConsumer: nextConsumer}
	b.reset()
	return b
}

func (bm *batchMetrics) export(ctx context.Context) error {
	return bm.nextConsumer.ConsumeMetrics(ctx, pdatautil.MetricsFromInternalMetrics(bm.metricData))
}

func (bm *batchMetrics) itemCount() uint32 {
	return bm.metricCount
}

// resets the current batchMetrics structure with zero/empty values.
func (bm *batchMetrics) reset() {
	bm.metricData = data.NewMetricData()
	bm.metricCount = 0
}

func (bm *batchMetrics) add(item interface{}) {
	md := pdatautil.MetricsToInternalMetrics(item.(pdata.Metrics))

	newMetricsCount := md.MetricCount()
	if newMetricsCount == 0 {
		return
	}
	bm.metricCount += uint32(newMetricsCount)
	md.ResourceMetrics().MoveAndAppendTo(bm.metricData.ResourceMetrics())
}
