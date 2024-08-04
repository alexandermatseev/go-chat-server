package repository

import (
	"context"

	"github.com/alexandermatseev/chat-server/internal/model"
)

// ChatRepository represents a chat repository.
type ChatRepository interface {
	Create(context context.Context, createChat *model.ChatCreate) (int64, error)
	Delete(context context.Context, id int64) error
}

// MessageRepository represents a message repository.
type MessageRepository interface {
	Send(context context.Context, createMessage *model.MessageCreate) (string, error)
}

// ContributorRepository represents a Contributor repository.
type ContributorRepository interface {
	CreateContributor(context context.Context, createContributor *model.ContributorCreate) error
	CreateContributors(context context.Context, createContributors *model.ContributorsCreate) error
	CheckContributorInChat(context context.Context, chatID int64, userID int64) (bool, error)
}
