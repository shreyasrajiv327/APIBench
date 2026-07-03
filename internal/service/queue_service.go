package service

import (
	"github.com/shreyasrajiv327/APIBench/internal/models"
	"github.com/shreyasrajiv327/APIBench/internal/repository"
	"fmt"
	"time"
)

type QueueService interface {
	Publish(payload []byte) (*models.Message, error)
	Poll() (*models.Message, error)
	GetMessage(messageID string) (*models.Message, error)
	Ack(messageID string) error
	Nack(messageID string) error
}

type queueService struct {
    repo repository.MessageRepository
}

func NewQueueService(repo repository.MessageRepository) QueueService {
	return &queueService{
		repo: repo,
	}
}

func (q *queueService) Publish(payload []byte) (*models.Message, error) {
	msg := &models.Message{
		ID:      generateID(),
		Payload: payload,
		Status:  models.StatusQueued,
		CreatedAt: time.Now(),

	}

	err := q.repo.Save(msg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

var counter = 0

func generateID() string {
	counter++
	return fmt.Sprintf("msg-%d", counter)
}

func (q *queueService) Poll() (*models.Message, error) {
	msg, err := q.repo.GetNext()
	if err != nil {
		return nil, err
	}

	msg.Status = models.StatusProcessing
	_ = q.repo.Update(msg)

	return msg, nil
}

func (q *queueService) Ack(messageID string) error {
	msg, err := q.repo.GetByID(messageID)
	if err != nil {
		return err
	}

	msg.Status = models.StatusAcked
	return q.repo.Update(msg)
}

func (q *queueService) Nack(messageID string) error {
	msg, err := q.repo.GetByID(messageID)
	if err != nil {
		return err
	}

	msg.Status = models.StatusNacked
	return q.repo.Update(msg)
}

func (q *queueService) GetMessage(messageID string) (*models.Message, error) {
	return q.repo.GetByID(messageID)
}