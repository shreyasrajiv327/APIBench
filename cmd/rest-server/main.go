package main

import (
	"github.com/gin-gonic/gin"

	"github.com/shreyasrajiv327/APIBench/api/rest"
	"github.com/shreyasrajiv327/APIBench/internal/repository"
	"github.com/shreyasrajiv327/APIBench/internal/service"
)

func main() {
	repo := repository.NewInMemoryRepository()
	queue := service.NewQueueService(repo)

	handler := rest.NewHandler(queue)

	router := gin.Default()

	rest.RegisterRoutes(router, handler)

	router.Run(":8080")
}