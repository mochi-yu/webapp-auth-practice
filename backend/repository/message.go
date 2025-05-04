package repository

import (
	"context"
	"time"

	"github.com/mochi-yu/webapp-auth-practice/ent"
	"github.com/mochi-yu/webapp-auth-practice/ent/message"
)

type IMessageRepository interface {
	Create(ctx context.Context, message *ent.Message) (*ent.Message, error)
	GetAll(ctx context.Context) ([]*ent.Message, error)
}

type MessageRepository struct {
	db *ent.Client
}

func NewMessageRepository(db *ent.Client) IMessageRepository {
	return &MessageRepository{db: db}
}

func (mr *MessageRepository) Create(ctx context.Context, message *ent.Message) (*ent.Message, error) {
	msg, err := mr.db.Message.Create().
		SetContent(message.Content).
		SetCreatedAt(time.Now()).
		SetAuthorName(message.AuthorName).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (mr *MessageRepository) GetAll(ctx context.Context) ([]*ent.Message, error) {
	msgs, err := mr.db.Message.Query().
		Order(ent.Desc(message.FieldCreatedAt)).All(ctx)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
