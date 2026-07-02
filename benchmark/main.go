package main

import (
	"fmt"
	"log"

	"github.com/shreyasrajiv327/APIBench/internal/benchmark"
)

func main() {
	runner := benchmark.NewRunner()

	fmt.Println("Running REST benchmark...")

	err := runner.RunREST("http://localhost:8080/messages", 1000)
	if err != nil {
		log.Fatal(err)
	}

	report := runner.Report()

	fmt.Println("\n========== RESULTS ==========")
	fmt.Printf("Total Requests : %d\n", report.TotalRequests)
	fmt.Printf("Average Latency: %v\n", report.Average)
	fmt.Printf("P95 Latency    : %v\n", report.P95)
	fmt.Printf("Throughput     : %.2f req/sec\n", report.Throughput)
}