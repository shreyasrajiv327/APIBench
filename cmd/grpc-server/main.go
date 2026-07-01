package main

import (
	"log"
	"net"

	apigrpc "github.com/shreyasrajiv327/APIBench/api/grpc"
	"github.com/shreyasrajiv327/APIBench/internal/repository"
	"github.com/shreyasrajiv327/APIBench/internal/service"
	pb "github.com/shreyasrajiv327/APIBench/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Create our existing QueueService
	repo := repository.NewInMemoryRepository()
	queue := service.NewQueueService(repo)

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register reflection so grpcurl can discover the service
	reflection.Register(grpcServer)

	// Register our QueueService implementation
	pb.RegisterQueueServiceServer(grpcServer, apigrpc.NewServer(queue))

	// Listen on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("gRPC Server running on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}