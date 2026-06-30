package main

import (
	"fmt"

	"github.com/shreyasrajiv327/APIBench/internal/repository"
	"github.com/shreyasrajiv327/APIBench/internal/service"
)

func main() {
	// init repository
	repo := repository.NewInMemoryRepository()

	// init service
	queue := service.NewQueueService(repo)

	// publish messages
	m1, _ := queue.Publish([]byte("hello"))
	m2, _ := queue.Publish([]byte("world"))

	fmt.Println("Published:", m1.ID, string(m1.Payload))
	fmt.Println("Published:", m2.ID, string(m2.Payload))

	// poll first message
	p1, err := queue.Poll()
	if err != nil {
		fmt.Println("Poll error:", err)
		return
	}
	fmt.Println("Polled:", p1.ID, string(p1.Payload))

	// ack first message
	_ = queue.Ack(p1.ID)

	// poll second message
	p2, err := queue.Poll()
	if err != nil {
		fmt.Println("Poll error:", err)
		return
	}
	fmt.Println("Polled:", p2.ID, string(p2.Payload))
}