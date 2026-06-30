package repository

import (
	"github.com/shreyasrajiv327/APIBench/internal/models"
	"errors"
	"sync"
)

type InMemoryRepository struct {
	mu       sync.RWMutex
	messages map[string]*models.Message
	queue    []*models.Message
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		messages: make(map[string]*models.Message),
		queue:    make([]*models.Message, 0),
	}
}

func (r *InMemoryRepository) Save(message *models.Message) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.messages[message.ID] = message
	r.queue = append(r.queue, message)

	return nil
}

func (r *InMemoryRepository) GetByID(id string) (*models.Message, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	message, exists := r.messages[id]
	if !exists {
		return nil, errors.New("message not found")
	}

	return message, nil
}

func (r *InMemoryRepository) GetNext() (*models.Message, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.queue) == 0 {
		return nil, errors.New("queue is empty")
	}

	message := r.queue[0]
	r.queue = r.queue[1:]

	return message, nil
}

func (r *InMemoryRepository) Update(message *models.Message) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.messages[message.ID]; !exists {
		return errors.New("message not found")
	}

	r.messages[message.ID] = message

	return nil
}