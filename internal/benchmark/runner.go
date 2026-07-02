package benchmark

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
	"context"

	"github.com/shreyasrajiv327/APIBench/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

func (r *Runner) RunGraphQL(url string, requests int) error {
	client := &http.Client{}

	query := map[string]string{
		"query": `mutation {
			publish(payload: "benchmark") {
				id
			}
		}`,
	}

	for i := 0; i < requests; i++ {
		data, _ := json.Marshal(query)

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

func (r *Runner) RunGRPC(address string, requests int) error {
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewQueueServiceClient(conn)

	for i := 0; i < requests; i++ {
		start := time.Now()

		_, err := client.Publish(
			context.Background(),
			&proto.PublishRequest{
				Payload: []byte("benchmark"),
			},
		)
		if err != nil {
			return err
		}

		r.metrics.Add(time.Since(start))
	}

	return nil
}

func (r *Runner) Report() Report {
	return r.metrics.Report()
}