package benchmark

import (
	"sort"
	"time"
)

type Report struct {
	TotalRequests int
	Average       time.Duration
	P95           time.Duration
	Throughput    float64
}

func (m *Metrics) Report() Report {
	latencies := make([]time.Duration, len(m.Latencies))
	copy(latencies, m.Latencies)

	sort.Slice(latencies, func(i, j int) bool {
		return latencies[i] < latencies[j]
	})

	avg := time.Duration(0)
	if m.TotalRequests > 0 {
		avg = m.TotalTime / time.Duration(m.TotalRequests)
	}

	p95 := time.Duration(0)
	if len(latencies) > 0 {
		index := int(float64(len(latencies)-1) * 0.95)
		p95 = latencies[index]
	}

	throughput := 0.0
	if m.TotalTime > 0 {
		throughput = float64(m.TotalRequests) / m.TotalTime.Seconds()
	}

	return Report{
		TotalRequests: m.TotalRequests,
		Average:       avg,
		P95:           p95,
		Throughput:    throughput,
	}
}