package analytics

import (
	"testing"

	"net/http/httptest"

	"time"

	"os"

	"github.com/alanfoster/configurable-proxy/clock"
	"github.com/stretchr/testify/assert"
)

func TestNewStats(t *testing.T) {
	mc := clock.NewMockClock(time.Unix(0, 0))
	stats := NewStats(mc)
	assert.Equal(t, len(stats.Requests), 0)
}

func TestStats_RecordStart(t *testing.T) {
	mc := clock.NewMockClock(time.Unix(0, 0))
	stats := NewStats(mc)

	r := httptest.NewRequest("GET", "http://example.com/foo", nil)
	id := stats.RecordStart(r)

	assert.Equal(t, 0, int(id))
	assert.Equal(t, []Request{
		Request{
			Start:      mc.Now(),
			URI:        "http://example.com/foo",
			Method:     "GET",
			Headers:    map[string]string{},
			Latency:    Latency{0},
			IsComplete: false,
		},
	}, stats.Requests)
}

func TestStats_RecordFinish(t *testing.T) {
	mc := clock.NewMockClock(time.Unix(0, int64(500*time.Millisecond)))
	stats := &Stats{
		Pid:   os.Getpid(),
		clock: mc,
		Requests: []Request{
			{
				Start:      time.Unix(0, 0),
				URI:        "http://example.com/foo",
				Method:     "GET",
				Headers:    map[string]string{},
				Latency:    Latency{0},
				IsComplete: false,
			},
		},
	}

	r := httptest.NewRequest("GET", "http://example.com/foo", nil)
	stats.RecordFinish(Id(0), r)

	assert.Equal(t, []Request{
		Request{
			Start:      time.Unix(0, 0),
			URI:        "http://example.com/foo",
			Method:     "GET",
			Headers:    map[string]string{},
			Latency:    Latency{500 * time.Millisecond},
			IsComplete: true,
		},
	}, stats.Requests)
}
