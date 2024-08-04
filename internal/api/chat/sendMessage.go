package chat

import (
	"context"

	"github.com/alexandermatseev/chat-server/internal/converter"
	cht "github.com/alexandermatseev/chat-server/pkg/chat_v1"
)

// SendMessage sends a new message
func (i *Implementation) SendMessage(ctx context.Context, req *cht.SendMessageRequest) (*cht.SendMessageResponse, error) {
	id, err := i.chatService.SendMessage(ctx, converter.ToMessageCreateFromcht(req.GetMessage()))
	if err != nil {
		return nil, err
	}
	return &cht.SendMessageResponse{
		Id:     id,
		ChatId: req.GetMessage().GetToChatId(),
	}, nil
}
