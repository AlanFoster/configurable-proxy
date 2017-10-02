package analytics

import (
	"encoding/json"
	"fmt"
	"time"
)

type Latency struct {
	time.Duration
}

func (l *Latency) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%.6f", l.Seconds()))
}

type Request struct {
	Start      time.Time         `json:"start"`
	URI        string            `json:"uri"`
	Method     string            `json:"method"`
	Headers    map[string]string `json:"headers"`
	Latency    Latency           `json:"latency_seconds"`
	IsComplete bool              `json:"is_complete"`
}
