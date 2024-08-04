package chat

import (
	"context"

	cht "github.com/alexandermatseev/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete deletes a chat
func (i *Implementation) Delete(ctx context.Context, req *cht.DeleteRequest) (*emptypb.Empty, error) {
	err := i.chatService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return nil, nil
}
