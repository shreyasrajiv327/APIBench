package rest

import "github.com/shreyasrajiv327/APIBench/internal/models"

type MessageResponse struct {
	ID        string `json:"id"`
	Payload   string `json:"payload"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
}

func ToMessageResponse(msg *models.Message) MessageResponse {
	return MessageResponse{
		ID:        msg.ID,
		Payload:   string(msg.Payload),
		Status:    string(msg.Status),
		CreatedAt: msg.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}