package message

import (
	"github.com/alexandermatseev/chat-server/internal/repository"
	"github.com/alexandermatseev/platform_common/pkg/db"
)

const (
	tableName = "messages"

	idColumn        = "id"
	chatIDColumn    = "chat_id"
	userIDColumn    = "user_id"
	textColumn      = "text"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

// NewRepository creates a new user repository.
func NewRepository(db db.Client) repository.MessageRepository {
	return &repo{db: db}
}
