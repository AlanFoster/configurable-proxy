package analytics

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"fmt"

	"github.com/stretchr/testify/assert"
)

var handlerTests = []struct {
	Description string
	Pid         int
	Requests    []Request
	Expected    string
}{
	{
		Description: "When there are no requests",
		Pid:         1337,
		Requests:    []Request{},
		Expected:    `{"pid":1337,"inflight_count":0,"requests":[]}` + "\n",
	},
	{
		Description: "When there are is one pending request",
		Pid:         1337,
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
		Expected: `{"pid":1337,"inflight_count":1,"requests":[{"start":"1970-01-01T01:00:00+01:00","uri":"http://example.com/foo","method":"GET","headers":{},"latency_seconds":"0.000000","is_complete":false}]}` + "\n",
	},
	{
		Description: "When there are is one complete request",
		Pid:         1337,
		Requests: []Request{
			{
				Start:      time.Unix(0, 0),
				URI:        "http://example.com/foo",
				Method:     "GET",
				Headers:    map[string]string{},
				Latency:    Latency{500 * time.Millisecond},
				IsComplete: true,
			},
		},
		Expected: `{"pid":1337,"inflight_count":0,"requests":[{"start":"1970-01-01T01:00:00+01:00","uri":"http://example.com/foo","method":"GET","headers":{},"latency_seconds":"0.500000","is_complete":true}]}` + "\n",
	},
}

func TestHandler_ServeHTTP(t *testing.T) {
	t.Parallel()

	for idx, test := range handlerTests {
		t.Run(fmt.Sprintf("%d %s", idx, test.Description), func(t *testing.T) {
			s := Stats{
				Pid:      test.Pid,
				Requests: test.Requests,
			}

			h := Handler{stats: &s}

			req, err := http.NewRequest("GET", "/stats", nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()

			h.ServeHTTP(rr, req)

			assert.Equal(t, test.Expected, rr.Body.String())
		})
	}
}
