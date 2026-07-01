package graphql

import "github.com/shreyasrajiv327/APIBench/internal/service"

type Resolver struct {
	Queue service.QueueService
}