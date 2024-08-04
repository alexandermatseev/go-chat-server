package chat

import (
	"github.com/alexandermatseev/chat-server/internal/client/authservice"
	"github.com/alexandermatseev/chat-server/internal/repository"
	"github.com/alexandermatseev/chat-server/internal/service"
	"github.com/alexandermatseev/platform_common/pkg/db"
)

type serv struct {
	chatRepository        repository.ChatRepository
	messageRepository     repository.MessageRepository
	ContributorRepository repository.ContributorRepository
	authServiceClient     authservice.AuthService
	txManager             db.TxManager
}

// NewService creates a new chat service.
func NewService(
	chatRepository repository.ChatRepository,
	messageRepository repository.MessageRepository,
	ContributorRepository repository.ContributorRepository,
	authServiceClient authservice.AuthService,
	txManager db.TxManager,
) service.ChatService {
	return &serv{
		chatRepository:        chatRepository,
		messageRepository:     messageRepository,
		ContributorRepository: ContributorRepository,
		authServiceClient:     authServiceClient,
		txManager:             txManager,
	}
}
