package analytics

import (
	"net/http"
	"os"
	"sync"

	"github.com/alanfoster/configurable-proxy/clock"
)

type Id int

type Stats struct {
	mu       sync.RWMutex
	Pid      int
	clock    clock.Clock
	Requests []Request
}

func NewStats(clock clock.Clock) *Stats {
	return &Stats{
		Pid:      os.Getpid(),
		clock:    clock,
		Requests: make([]Request, 0, 10),
	}
}

func (stats *Stats) Reset() {
	stats.mu.Lock()
	defer stats.mu.Unlock()
	stats.Requests = make([]Request, 0, 10)
}

func (stats *Stats) RecordStart(request *http.Request) Id {
	stats.mu.Lock()
	defer stats.mu.Unlock()

	headers := make(map[string]string)

	for header, value := range headers {
		headers[header] = value
	}

	stats.Requests = append(stats.Requests, Request{
		URI:        request.RequestURI,
		Method:     request.Method,
		Headers:    headers,
		Start:      stats.clock.Now(),
		IsComplete: false,
	})

	return Id(len(stats.Requests) - 1)
}

func (stats *Stats) RecordFinish(id Id, r *http.Request) {
	stats.mu.Lock()
	defer stats.mu.Unlock()

	request := stats.Requests[id]
	request.IsComplete = true
	request.Latency = Latency{stats.clock.Now().Sub(request.Start)}
	stats.Requests[id] = request
}
