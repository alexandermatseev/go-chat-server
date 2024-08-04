package chat

import (
	"context"
	"log"

	"github.com/alexandermatseev/chat-server/internal/converter"
	cht "github.com/alexandermatseev/chat-server/pkg/chat_v1"
)

// Create creates a new chat
func (i *Implementation) Create(ctx context.Context, req *cht.CreateRequest) (*cht.CreateResponse, error) {
	id, err := i.chatService.Create(ctx, converter.ToChatCreateFromcht(req.GetChat()), converter.ToContributorsCreateFromcht(req.GetUserIds()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted chat with id: %d", id)

	return &cht.CreateResponse{
		Id: id,
	}, nil
}
