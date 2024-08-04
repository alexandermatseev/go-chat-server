package chat

import (
	"context"

	"github.com/pkg/errors"

	"github.com/alexandermatseev/chat-server/internal/model"
)

func (s *serv) SendMessage(ctx context.Context, createMessage *model.MessageCreate) (string, error) {
	isExists, err := s.ContributorRepository.CheckContributorInChat(ctx, createMessage.Info.ChatID, createMessage.Info.UserID)
	if !isExists || err != nil {
		return "", errors.New("user is not a Contributor of the chat")
	}
	id, err := s.messageRepository.Send(ctx, createMessage)
	if err != nil {
		return "", err
	}
	return id, nil
}
