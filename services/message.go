package services

import (
	"context"

	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/repositories"
)

type IMessageService interface {
	Create(input model.NewMessage, ctx context.Context) (*model.Messages, error)
	GetMessages(from string, to string, ctx context.Context) ([]*model.Messages, error)
}

type MessageService struct {
	MessageRepository *repositories.MessageRepository
}

func NewMessageService(messageRepository *repositories.MessageRepository) *MessageService {
	return &MessageService{
		MessageRepository: messageRepository,
	}
}

func (r *MessageService) Create(input model.NewMessage, ctx context.Context) (*model.Messages, error) {
	return r.MessageRepository.Create(input, ctx)
}

func (r *MessageService) GetMessages(from string, to string, ctx context.Context) ([]*model.Messages, error) {
	return r.MessageRepository.GetMessages(from, to, ctx)
}
