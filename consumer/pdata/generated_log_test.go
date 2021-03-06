// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "go run cmd/pdatagen/main.go".

package pdata

import (
	"testing"

	"github.com/stretchr/testify/assert"

	logsproto "go.opentelemetry.io/collector/internal/data/opentelemetry-proto-gen/logs/v1"
)

func TestResourceLogsSlice(t *testing.T) {
	es := NewResourceLogsSlice()
	assert.EqualValues(t, 0, es.Len())
	es = newResourceLogsSlice(&[]*logsproto.ResourceLogs{})
	assert.EqualValues(t, 0, es.Len())

	es.Resize(7)
	emptyVal := NewResourceLogs()
	emptyVal.InitEmpty()
	testVal := generateTestResourceLogs()
	assert.EqualValues(t, 7, es.Len())
	for i := 0; i < es.Len(); i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
		fillTestResourceLogs(es.At(i))
		assert.EqualValues(t, testVal, es.At(i))
	}
}

func TestResourceLogsSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestResourceLogsSlice()
	dest := NewResourceLogsSlice()
	src := generateTestResourceLogsSlice()
	src.MoveAndAppendTo(dest)
	assert.EqualValues(t, generateTestResourceLogsSlice(), dest)
	assert.EqualValues(t, 0, src.Len())
	assert.EqualValues(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.EqualValues(t, generateTestResourceLogsSlice(), dest)
	assert.EqualValues(t, 0, src.Len())
	assert.EqualValues(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestResourceLogsSlice().MoveAndAppendTo(dest)
	assert.EqualValues(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.EqualValues(t, expectedSlice.At(i), dest.At(i))
		assert.EqualValues(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestResourceLogsSlice_CopyTo(t *testing.T) {
	dest := NewResourceLogsSlice()
	// Test CopyTo to empty
	NewResourceLogsSlice().CopyTo(dest)
	assert.EqualValues(t, NewResourceLogsSlice(), dest)

	// Test CopyTo larger slice
	generateTestResourceLogsSlice().CopyTo(dest)
	assert.EqualValues(t, generateTestResourceLogsSlice(), dest)

	// Test CopyTo same size slice
	generateTestResourceLogsSlice().CopyTo(dest)
	assert.EqualValues(t, generateTestResourceLogsSlice(), dest)
}

func TestResourceLogsSlice_Resize(t *testing.T) {
	es := generateTestResourceLogsSlice()
	emptyVal := NewResourceLogs()
	emptyVal.InitEmpty()
	// Test Resize less elements.
	const resizeSmallLen = 4
	expectedEs := make(map[*logsproto.ResourceLogs]bool, resizeSmallLen)
	for i := 0; i < resizeSmallLen; i++ {
		expectedEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, resizeSmallLen, len(expectedEs))
	es.Resize(resizeSmallLen)
	assert.EqualValues(t, resizeSmallLen, es.Len())
	foundEs := make(map[*logsproto.ResourceLogs]bool, resizeSmallLen)
	for i := 0; i < es.Len(); i++ {
		foundEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)

	// Test Resize more elements.
	const resizeLargeLen = 7
	oldLen := es.Len()
	expectedEs = make(map[*logsproto.ResourceLogs]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, oldLen, len(expectedEs))
	es.Resize(resizeLargeLen)
	assert.EqualValues(t, resizeLargeLen, es.Len())
	foundEs = make(map[*logsproto.ResourceLogs]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)
	for i := oldLen; i < resizeLargeLen; i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
	}

	// Test Resize 0 elements.
	es.Resize(0)
	assert.EqualValues(t, NewResourceLogsSlice(), es)
}

func TestResourceLogsSlice_Append(t *testing.T) {
	es := generateTestResourceLogsSlice()
	emptyVal := NewResourceLogs()
	emptyVal.InitEmpty()

	es.Append(&emptyVal)
	assert.EqualValues(t, *(es.At(7)).orig, *emptyVal.orig)

	emptyVal2:= NewResourceLogs()
	emptyVal2.InitEmpty()

	es.Append(&emptyVal2)
	assert.EqualValues(t, *(es.At(8)).orig, *emptyVal2.orig)

	assert.Equal(t, 9, es.Len())
}

func TestResourceLogs_InitEmpty(t *testing.T) {
	ms := NewResourceLogs()
	assert.True(t, ms.IsNil())
	ms.InitEmpty()
	assert.False(t, ms.IsNil())
}

func TestResourceLogs_CopyTo(t *testing.T) {
	ms := NewResourceLogs()
	NewResourceLogs().CopyTo(ms)
	assert.True(t, ms.IsNil())
	generateTestResourceLogs().CopyTo(ms)
	assert.EqualValues(t, generateTestResourceLogs(), ms)
}

func TestResourceLogs_Resource(t *testing.T) {
	ms := NewResourceLogs()
	ms.InitEmpty()
	assert.EqualValues(t, true, ms.Resource().IsNil())
	ms.Resource().InitEmpty()
	assert.EqualValues(t, false, ms.Resource().IsNil())
	fillTestResource(ms.Resource())
	assert.EqualValues(t, generateTestResource(), ms.Resource())
}

func TestResourceLogs_Logs(t *testing.T) {
	ms := NewResourceLogs()
	ms.InitEmpty()
	assert.EqualValues(t, NewLogSlice(), ms.Logs())
	fillTestLogSlice(ms.Logs())
	testValLogs := generateTestLogSlice()
	assert.EqualValues(t, testValLogs, ms.Logs())
}

func TestLogSlice(t *testing.T) {
	es := NewLogSlice()
	assert.EqualValues(t, 0, es.Len())
	es = newLogSlice(&[]*logsproto.LogRecord{})
	assert.EqualValues(t, 0, es.Len())

	es.Resize(7)
	emptyVal := NewLogRecord()
	emptyVal.InitEmpty()
	testVal := generateTestLogRecord()
	assert.EqualValues(t, 7, es.Len())
	for i := 0; i < es.Len(); i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
		fillTestLogRecord(es.At(i))
		assert.EqualValues(t, testVal, es.At(i))
	}
}

func TestLogSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestLogSlice()
	dest := NewLogSlice()
	src := generateTestLogSlice()
	src.MoveAndAppendTo(dest)
	assert.EqualValues(t, generateTestLogSlice(), dest)
	assert.EqualValues(t, 0, src.Len())
	assert.EqualValues(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.EqualValues(t, generateTestLogSlice(), dest)
	assert.EqualValues(t, 0, src.Len())
	assert.EqualValues(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestLogSlice().MoveAndAppendTo(dest)
	assert.EqualValues(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.EqualValues(t, expectedSlice.At(i), dest.At(i))
		assert.EqualValues(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestLogSlice_CopyTo(t *testing.T) {
	dest := NewLogSlice()
	// Test CopyTo to empty
	NewLogSlice().CopyTo(dest)
	assert.EqualValues(t, NewLogSlice(), dest)

	// Test CopyTo larger slice
	generateTestLogSlice().CopyTo(dest)
	assert.EqualValues(t, generateTestLogSlice(), dest)

	// Test CopyTo same size slice
	generateTestLogSlice().CopyTo(dest)
	assert.EqualValues(t, generateTestLogSlice(), dest)
}

func TestLogSlice_Resize(t *testing.T) {
	es := generateTestLogSlice()
	emptyVal := NewLogRecord()
	emptyVal.InitEmpty()
	// Test Resize less elements.
	const resizeSmallLen = 4
	expectedEs := make(map[*logsproto.LogRecord]bool, resizeSmallLen)
	for i := 0; i < resizeSmallLen; i++ {
		expectedEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, resizeSmallLen, len(expectedEs))
	es.Resize(resizeSmallLen)
	assert.EqualValues(t, resizeSmallLen, es.Len())
	foundEs := make(map[*logsproto.LogRecord]bool, resizeSmallLen)
	for i := 0; i < es.Len(); i++ {
		foundEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)

	// Test Resize more elements.
	const resizeLargeLen = 7
	oldLen := es.Len()
	expectedEs = make(map[*logsproto.LogRecord]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, oldLen, len(expectedEs))
	es.Resize(resizeLargeLen)
	assert.EqualValues(t, resizeLargeLen, es.Len())
	foundEs = make(map[*logsproto.LogRecord]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[*(es.At(i).orig)] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)
	for i := oldLen; i < resizeLargeLen; i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
	}

	// Test Resize 0 elements.
	es.Resize(0)
	assert.EqualValues(t, NewLogSlice(), es)
}

func TestLogSlice_Append(t *testing.T) {
	es := generateTestLogSlice()
	emptyVal := NewLogRecord()
	emptyVal.InitEmpty()

	es.Append(&emptyVal)
	assert.EqualValues(t, *(es.At(7)).orig, *emptyVal.orig)

	emptyVal2:= NewLogRecord()
	emptyVal2.InitEmpty()

	es.Append(&emptyVal2)
	assert.EqualValues(t, *(es.At(8)).orig, *emptyVal2.orig)

	assert.Equal(t, 9, es.Len())
}

func TestLogRecord_InitEmpty(t *testing.T) {
	ms := NewLogRecord()
	assert.True(t, ms.IsNil())
	ms.InitEmpty()
	assert.False(t, ms.IsNil())
}

func TestLogRecord_CopyTo(t *testing.T) {
	ms := NewLogRecord()
	NewLogRecord().CopyTo(ms)
	assert.True(t, ms.IsNil())
	generateTestLogRecord().CopyTo(ms)
	assert.EqualValues(t, generateTestLogRecord(), ms)
}

func TestLogRecord_Timestamp(t *testing.T) {
	ms := NewLogRecord()
	ms.InitEmpty()
	assert.EqualValues(t, TimestampUnixNano(0), ms.Timestamp())
	testValTimestamp := TimestampUnixNano(1234567890)
	ms.SetTimestamp(testValTimestamp)
	assert.EqualValues(t, testValTimestamp, ms.Timestamp())
}

func TestLogRecord_TraceID(t *testing.T) {
	ms := NewLogRecord()
	ms.InitEmpty()
	assert.EqualValues(t, NewTraceID(nil), ms.TraceID())
	testValTraceID := NewTraceID([]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1})
	ms.SetTraceID(testValTraceID)
	assert.EqualValues(t, testValTraceID, ms.TraceID())
}

func TestLogRecord_SpanID(t *testing.T) {
	ms := NewLogRecord()
	ms.InitEmpty()
	assert.EqualValues(t, NewSpanID(nil), ms.SpanID())
	testValSpanID := NewSpanID([]byte{1, 2, 3, 4, 5, 6, 7, 8})
	ms.SetSpanID(testValSpanID)
	assert.EqualValues(t, testValSpanID, ms.SpanID())
}

func TestLogRecord_Flags(t *testing.T) {
	ms := NewLogRecord()
	ms.InitEmpty()
	assert.EqualValues(t, uint32(0), ms.Flags())
	testValFlags := uint32(0x01)
	ms.SetFlags(testValFlags)
	assert.EqualValues(t, testValFlags, ms.Flags())
}

func TestLogRecord_SeverityText(t *testing.T) {
	ms := NewLogRecord()
	ms.InitEmpty()
	assert.EqualValues(t, "", ms.SeverityText())
	testValSeverityText := "INFO"
	ms.SetSeverityText(testValSeverityText)
	assert.EqualValues(t, testValSeverityText, ms.SeverityText())
}

func TestLogRecord_SeverityNumber(t *testing.T) {
	ms := NewLogRecord()
	ms.InitEmpty()
	assert.EqualValues(t, logsproto.SeverityNumber_UNDEFINED_SEVERITY_NUMBER, ms.SeverityNumber())
	testValSeverityNumber := logsproto.SeverityNumber_INFO
	ms.SetSeverityNumber(testValSeverityNumber)
	assert.EqualValues(t, testValSeverityNumber, ms.SeverityNumber())
}

func TestLogRecord_ShortName(t *testing.T) {
	ms := NewLogRecord()
	ms.InitEmpty()
	assert.EqualValues(t, "", ms.ShortName())
	testValShortName := "test_name"
	ms.SetShortName(testValShortName)
	assert.EqualValues(t, testValShortName, ms.ShortName())
}

func TestLogRecord_Body(t *testing.T) {
	ms := NewLogRecord()
	ms.InitEmpty()
	assert.EqualValues(t, "", ms.Body())
	testValBody := "test log message"
	ms.SetBody(testValBody)
	assert.EqualValues(t, testValBody, ms.Body())
}

func TestLogRecord_Attributes(t *testing.T) {
	ms := NewLogRecord()
	ms.InitEmpty()
	assert.EqualValues(t, NewAttributeMap(), ms.Attributes())
	fillTestAttributeMap(ms.Attributes())
	testValAttributes := generateTestAttributeMap()
	assert.EqualValues(t, testValAttributes, ms.Attributes())
}

func TestLogRecord_DroppedAttributesCount(t *testing.T) {
	ms := NewLogRecord()
	ms.InitEmpty()
	assert.EqualValues(t, uint32(0), ms.DroppedAttributesCount())
	testValDroppedAttributesCount := uint32(17)
	ms.SetDroppedAttributesCount(testValDroppedAttributesCount)
	assert.EqualValues(t, testValDroppedAttributesCount, ms.DroppedAttributesCount())
}

func generateTestResourceLogsSlice() ResourceLogsSlice {
	tv := NewResourceLogsSlice()
	fillTestResourceLogsSlice(tv)
	return tv
}

func fillTestResourceLogsSlice(tv ResourceLogsSlice) {
	tv.Resize(7)
	for i := 0; i < tv.Len(); i++ {
		fillTestResourceLogs(tv.At(i))
	}
}

func generateTestResourceLogs() ResourceLogs {
	tv := NewResourceLogs()
	tv.InitEmpty()
	fillTestResourceLogs(tv)
	return tv
}

func fillTestResourceLogs(tv ResourceLogs) {
	tv.Resource().InitEmpty()
	fillTestResource(tv.Resource())
	fillTestLogSlice(tv.Logs())
}

func generateTestLogSlice() LogSlice {
	tv := NewLogSlice()
	fillTestLogSlice(tv)
	return tv
}

func fillTestLogSlice(tv LogSlice) {
	tv.Resize(7)
	for i := 0; i < tv.Len(); i++ {
		fillTestLogRecord(tv.At(i))
	}
}

func generateTestLogRecord() LogRecord {
	tv := NewLogRecord()
	tv.InitEmpty()
	fillTestLogRecord(tv)
	return tv
}

func fillTestLogRecord(tv LogRecord) {
	tv.SetTimestamp(TimestampUnixNano(1234567890))
	tv.SetTraceID(NewTraceID([]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1}))
	tv.SetSpanID(NewSpanID([]byte{1, 2, 3, 4, 5, 6, 7, 8}))
	tv.SetFlags(uint32(0x01))
	tv.SetSeverityText("INFO")
	tv.SetSeverityNumber(logsproto.SeverityNumber_INFO)
	tv.SetShortName("test_name")
	tv.SetBody("test log message")
	fillTestAttributeMap(tv.Attributes())
	tv.SetDroppedAttributesCount(uint32(17))
}
