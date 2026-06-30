package service

import "APIbench/internal/models"

type QueueService interface{
	Publish(payload []byte) (*models.Message, error)
	Poll() (*models.Message, error)
	Ack(messageID string) error
	Nack(messageID string) error
} 

