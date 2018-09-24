package analytics

import (
	"net/http"
)

type Middleware struct {
	stats *Stats
}

func NewMiddleware(stats *Stats) Middleware {
	return Middleware{
		stats: stats,
	}
}

func (m *Middleware) Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := m.stats.RecordStart(r)
		h.ServeHTTP(w, r)
		m.stats.RecordFinish(id, r)
	})
}
