package contributor

import (
	"github.com/alexandermatseev/chat-server/internal/repository"
	"github.com/alexandermatseev/platform_common/pkg/db"
)

const (
	tableName = "chatContributors"

	idColumn        = "id"
	chatIDColumn    = "chat_id"
	userIDColumn    = "user_id"
	createdAtColumn = "created_at"
)

type repo struct {
	db db.Client
}

// NewRepository creates a new user repository.
func NewRepository(db db.Client) repository.ContributorRepository {
	return &repo{db: db}
}
