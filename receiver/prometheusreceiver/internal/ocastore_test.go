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

package internal

import (
	"context"
	"testing"

	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/scrape"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOcaStore(t *testing.T) {

	o := NewOcaStore(context.Background(), nil, nil, nil, false, "prometheus")

	_, err := o.Appender()
	require.Error(t, err)

	o.SetScrapeManager(nil)
	_, err = o.Appender()
	require.Error(t, err, "Expecting error when ScrapeManager is not set")

	o.SetScrapeManager(&scrape.Manager{})

	app, err := o.Appender()
	require.NotNil(t, app, "Expecting app, but got error %v\n", err)

	_ = o.Close()

	app, err = o.Appender()
	assert.Equal(t, noop, app)
	assert.NoError(t, err)
}

func TestNoopAppender(t *testing.T) {
	if _, err := noop.Add(labels.FromStrings("t", "v"), 1, 1); err == nil {
		t.Error("expecting error from Add method of noopApender")
	}
	if _, err := noop.Add(labels.FromStrings("t", "v"), 1, 1); err == nil {
		t.Error("expecting error from Add method of noopApender")
	}

	if err := noop.AddFast(labels.FromStrings("t", "v"), 0, 1, 1); err == nil {
		t.Error("expecting error from AddFast method of noopApender")
	}

	if err := noop.Commit(); err == nil {
		t.Error("expecting error from Commit method of noopApender")
	}

	if err := noop.Rollback(); err != nil {
		t.Error("expecting no error from Rollback method of noopApender")
	}

}
