package main

import (
	"log"
	"net/http"

	api "github.com/shreyasrajiv327/APIBench/api/graphql"
	"github.com/shreyasrajiv327/APIBench/internal/repository"
	"github.com/shreyasrajiv327/APIBench/internal/service"
)

func main() {

	repo := repository.NewInMemoryRepository()
	queue := service.NewQueueService(repo)

	resolver := &api.Resolver{
		Queue: queue,
	}

	http.Handle("/graphql", api.NewHandler(resolver))

	log.Println("GraphQL Server running on :8081")

	log.Fatal(http.ListenAndServe(":8081", nil))
}