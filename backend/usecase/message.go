package usecase

import (
	"context"

	"github.com/mochi-yu/webapp-auth-practice/ent"
	"github.com/mochi-yu/webapp-auth-practice/repository"
)

type IMessageUsecase interface {
	Create(ctx context.Context, message *ent.Message) (*ent.Message, error)
	GetAll(ctx context.Context) ([]*ent.Message, error)
}

type MessageUsecase struct {
	mr repository.IMessageRepository
}

func NewMessageUsecase(mr repository.IMessageRepository) IMessageUsecase {
	return &MessageUsecase{mr: mr}
}

func (mu *MessageUsecase) Create(ctx context.Context, message *ent.Message) (*ent.Message, error) {
	msg, err := mu.mr.Create(ctx, message)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (mu *MessageUsecase) GetAll(ctx context.Context) ([]*ent.Message, error) {
	msgs, err := mu.mr.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
