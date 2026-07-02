package benchmark

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type Runner struct {
	metrics *Metrics
}

func NewRunner() *Runner {
	return &Runner{
		metrics: &Metrics{},
	}
}
func (r *Runner) RunREST(url string, requests int) error {
	client := &http.Client{}

	for i := 0; i < requests; i++ {

		body := map[string]string{
			"payload": "benchmark",
		}

		data, _ := json.Marshal(body)

		start := time.Now()

		resp, err := client.Post(
			url,
			"application/json",
			bytes.NewBuffer(data),
		)
		if err != nil {
			return err
		}

		resp.Body.Close()

		r.metrics.Add(time.Since(start))
	}

	return nil
}

func (r *Runner) Report() Report {
	return r.metrics.Report()
}