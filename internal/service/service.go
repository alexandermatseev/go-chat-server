package service

import (
	"context"

	"github.com/alexandermatseev/chat-server/internal/model"
)

// ChatService represents a chat service.
type ChatService interface {
	Create(ctx context.Context, createChat *model.ChatCreate, createContributors *model.ContributorsCreate) (int64, error)
	SendMessage(ctx context.Context, createMessage *model.MessageCreate) (string, error)
	Delete(ctx context.Context, id int64) error
}
