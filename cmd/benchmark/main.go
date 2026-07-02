package main

import (
	"fmt"
	"log"

	"github.com/shreyasrajiv327/APIBench/internal/benchmark"
)

func printReport(name string, report benchmark.Report) {
	fmt.Printf("\n===== %s =====\n", name)
	fmt.Printf("Total Requests : %d\n", report.TotalRequests)
	fmt.Printf("Min Latency    : %v\n", report.Min)
	fmt.Printf("Median Latency : %v\n", report.Median)
	fmt.Printf("Average Latency: %v\n", report.Average)
	fmt.Printf("P95 Latency    : %v\n", report.P95)
	fmt.Printf("P99 Latency    : %v\n", report.P99)
	fmt.Printf("Max Latency    : %v\n", report.Max)
	fmt.Printf("Throughput     : %.2f req/sec\n", report.Throughput)
}

func main() {

	// ---------------- REST ----------------
	restRunner := benchmark.NewRunner()

	fmt.Println("Running REST benchmark...")

	if err := restRunner.RunREST("http://localhost:8080/messages", 1000); err != nil {
		log.Fatal(err)
	}

	printReport("REST", restRunner.Report())

	// ---------------- GraphQL ----------------
	graphqlRunner := benchmark.NewRunner()

	fmt.Println("\nRunning GraphQL benchmark...")

	if err := graphqlRunner.RunGraphQL("http://localhost:8081/graphql", 1000); err != nil {
		log.Fatal(err)
	}

	printReport("GraphQL", graphqlRunner.Report())

	// ---------------- gRPC ----------------
	grpcRunner := benchmark.NewRunner()

	fmt.Println("\nRunning gRPC benchmark...")

	if err := grpcRunner.RunGRPC("localhost:50051", 1000); err != nil {
		log.Fatal(err)
	}

	printReport("gRPC", grpcRunner.Report())
}