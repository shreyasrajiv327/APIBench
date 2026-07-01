package graphql

import (
	"time"

	"github.com/shreyasrajiv327/APIBench/internal/models"
)

type MessageResponse struct {
	ID        string
	Payload   string
	Status    string
	CreatedAt string
}

func ToMessageResponse(msg *models.Message) MessageResponse {
	return MessageResponse{
		ID:        msg.ID,
		Payload:   string(msg.Payload),
		Status:    string(msg.Status),
		CreatedAt: msg.CreatedAt.Format(time.RFC3339),
	}
}