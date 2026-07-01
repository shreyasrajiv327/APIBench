package grpc

import (
	"context"

	"github.com/shreyasrajiv327/APIBench/internal/service"
	pb "github.com/shreyasrajiv327/APIBench/proto"
)

type Server struct {
	pb.UnimplementedQueueServiceServer
	queue service.QueueService
}

func NewServer(queue service.QueueService) *Server {
	return &Server{
		queue: queue,
	}
}

func (s *Server) Publish(ctx context.Context, req *pb.PublishRequest) (*pb.Message, error) {
	msg, err := s.queue.Publish(req.Payload)
	if err != nil {
		return nil, err
	}

	return ToProtoMessage(msg), nil
}

func (s *Server) Poll(ctx context.Context, req *pb.Empty) (*pb.Message, error) {
	msg, err := s.queue.Poll()
	if err != nil {
		return nil, err
	}

	return ToProtoMessage(msg), nil
}

func (s *Server) Ack(ctx context.Context, req *pb.AckRequest) (*pb.Empty, error) {
	if err := s.queue.Ack(req.Id); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s *Server) Nack(ctx context.Context, req *pb.AckRequest) (*pb.Empty, error) {
	if err := s.queue.Nack(req.Id); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}