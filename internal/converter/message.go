package converter

import (
	"github.com/alexandermatseev/chat-server/internal/model"
	cht "github.com/alexandermatseev/chat-server/pkg/chat_v1"
)

// ToMessageCreateFromcht converts cht.MessageCreate to model.MessageCreate
func ToMessageCreateFromcht(messageCreate *cht.MessageCreate) *model.MessageCreate {
	return &model.MessageCreate{
		Info: model.MessageInfo{
			ChatID: messageCreate.ToChatId,
			UserID: messageCreate.FromUserId,
			Text:   messageCreate.Text,
		},
	}
}
