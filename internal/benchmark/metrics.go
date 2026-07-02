package benchmark

import "time"

type Metrics struct {
	TotalRequests int
	TotalTime     time.Duration
	Latencies     []time.Duration
}

func (m *Metrics) Add(latency time.Duration) {
	m.TotalRequests++
	m.TotalTime += latency
	m.Latencies = append(m.Latencies, latency)
}