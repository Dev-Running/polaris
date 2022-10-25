package repositories

import (
	"context"

	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
	"github.com/laurentino14/user/prisma/connect"
)

type IMessageRepository interface {
	Create(input model.NewMessage, ctx context.Context) (*model.Messages, error)
	GetMessages(from string, to string, ctx context.Context) ([]*model.Messages, error)
}

type MessageRepository struct {
	DB *connect.DB
}

func NewMessageRepository(db *connect.DB) *MessageRepository {
	return &MessageRepository{DB: db}
}

func (r *MessageRepository) Create(input model.NewMessage, ctx context.Context) (*model.Messages, error) {
	message, err := r.DB.Client.Messages.CreateOne(
		prisma.Messages.Message.Set(input.Message),
		prisma.Messages.From.Set(input.From),
		prisma.Messages.To.Set(input.To),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	messages := &model.Messages{
		ID:        message.ID,
		Message:   message.Message,
		From:      message.From,
		To:        message.To,
		CreatedAt: message.CreatedAt.String(),
		User:      nil,
		UserID:    nil,
	}
	return messages, nil
}

func (r *MessageRepository) GetMessages(from string, to string, ctx context.Context) ([]*model.Messages, error) {
	messages, err := r.DB.Client.Messages.FindMany(
		prisma.Messages.From.Equals(from),
		prisma.Messages.To.Equals(to),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	messages1, err := r.DB.Client.Messages.FindMany(
		prisma.Messages.From.Equals(to),
		prisma.Messages.To.Equals(from),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	m := []*model.Messages{}
	for _, message := range messages {
		m = append(m, &model.Messages{
			ID:        message.ID,
			Message:   message.Message,
			From:      message.From,
			To:        message.To,
			CreatedAt: message.CreatedAt.String(),
			User:      nil,
			UserID:    nil,
		})
	}

	for _, message := range messages1 {
		m = append(m, &model.Messages{
			ID:        message.ID,
			Message:   message.Message,
			From:      message.From,
			To:        message.To,
			CreatedAt: message.CreatedAt.String(),
			User:      nil,
			UserID:    nil,
		})
	}
	return m, nil
}
