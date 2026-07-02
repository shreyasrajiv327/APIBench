package benchmark

import (
	"math"
	"sort"
	"time"
)

type Report struct {
	TotalRequests int
	Min           time.Duration
	Median        time.Duration
	Average       time.Duration
	P95           time.Duration
	P99           time.Duration
	Max           time.Duration
	Throughput    float64
}

func (m *Metrics) Report() Report {
	if len(m.Latencies) == 0 {
		return Report{}
	}

	latencies := make([]time.Duration, len(m.Latencies))
	copy(latencies, m.Latencies)

	sort.Slice(latencies, func(i, j int) bool {
		return latencies[i] < latencies[j]
	})

	avg := time.Duration(0)
	if m.TotalRequests > 0 {
		avg = m.TotalTime / time.Duration(m.TotalRequests)
	}

	p95Index := int(math.Ceil(0.95*float64(len(latencies)))) - 1
	if p95Index < 0 {
		p95Index = 0
	}
	if p95Index >= len(latencies) {
		p95Index = len(latencies) - 1
	}

	p99Index := int(math.Ceil(0.99*float64(len(latencies)))) - 1
	if p99Index < 0 {
		p99Index = 0
	}
	if p99Index >= len(latencies) {
		p99Index = len(latencies) - 1
	}

	throughput := 0.0
	if m.TotalTime > 0 {
		throughput = float64(m.TotalRequests) / m.TotalTime.Seconds()
	}

	return Report{
		TotalRequests: m.TotalRequests,
		Min:           latencies[0],
		Median:        latencies[len(latencies)/2],
		Average:       avg,
		P95:           latencies[p95Index],
		P99:           latencies[p99Index],
		Max:           latencies[len(latencies)-1],
		Throughput:    throughput,
	}
}