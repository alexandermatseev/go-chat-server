package chat

import (
	"github.com/alexandermatseev/chat-server/internal/service"
	cht "github.com/alexandermatseev/chat-server/pkg/chat_v1"
)

// Implementation represents a chat API implementation.
type Implementation struct {
	cht.UnimplementedChatV1Server
	chatService service.ChatService
}

// NewImplementation creates a new chat API implementation.
func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
