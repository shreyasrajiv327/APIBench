package grpc

import (
	"time"

	pb "github.com/shreyasrajiv327/APIBench/proto"
	"github.com/shreyasrajiv327/APIBench/internal/models"
)

func ToProtoMessage(msg *models.Message) *pb.Message {
	return &pb.Message{
		Id:        msg.ID,
		Payload:   msg.Payload,
		Status:    string(msg.Status),
		CreatedAt: msg.CreatedAt.Format(time.RFC3339),
	}
}