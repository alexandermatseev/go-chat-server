package converter

import (
	"github.com/alexandermatseev/chat-server/internal/model"
	cht "github.com/alexandermatseev/chat-server/pkg/chat_v1"
)

// ToChatCreateFromcht converts cht.ChatCreate to model.ChatCreate
func ToChatCreateFromcht(chatCreate *cht.ChatCreate) *model.ChatCreate {
	return &model.ChatCreate{
		Name: chatCreate.Name,
	}
}
