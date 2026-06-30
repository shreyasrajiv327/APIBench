package repository

import "github.com/shreyasrajiv327/APIBench/internal/models"

type MessageRepository interface {
	Save(message *models.Message) error
	GetByID(id string) (*models.Message, error)
	GetNext() (*models.Message, error)
	Update(message *models.Message) error
}