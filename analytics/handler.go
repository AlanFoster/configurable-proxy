package analytics

import (
	"encoding/json"
	"net/http"
)

type Snapshot struct {
	Pid           int       `json:"pid"`
	InflightCount int       `json:"inflight_count"`
	Requests      []Request `json:"requests"`
}

func (stats *Stats) Snapshot() *Snapshot {
	stats.mu.RLock()
	defer stats.mu.RUnlock()

	inflightCount := 0
	for _, request := range stats.Requests {
		if !request.IsComplete {
			inflightCount++
		}
	}

	return &Snapshot{
		Pid:           stats.Pid,
		InflightCount: inflightCount,
		Requests:      stats.Requests,
	}
}

type Handler struct {
	stats *Stats
}

func NewHandler(stats *Stats) *Handler {
	return &Handler{stats: stats}
}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	snapshot := handler.stats.Snapshot()

	if err := json.NewEncoder(w).Encode(snapshot); err != nil {
		panic(err)
	}
}
