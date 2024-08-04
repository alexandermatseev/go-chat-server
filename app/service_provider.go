package app

import (
	"context"
	"log"

	"github.com/alexandermatseev/platform_common/pkg/closer"
	"github.com/alexandermatseev/platform_common/pkg/db"
	"github.com/alexandermatseev/platform_common/pkg/db/pg"
	"github.com/alexandermatseev/platform_common/pkg/db/transaction"

	"github.com/alexandermatseev/chat-server/internal/client/authservice"
	authService2 "github.com/alexandermatseev/chat-server/internal/client/authservice/authservice"

	"github.com/alexandermatseev/chat-server/internal/api/chat"
	"github.com/alexandermatseev/chat-server/internal/config"
	"github.com/alexandermatseev/chat-server/internal/repository"
	chatRepository "github.com/alexandermatseev/chat-server/internal/repository/chat"
	ContributorRepository "github.com/alexandermatseev/chat-server/internal/repository/contributor"
	messageRepository "github.com/alexandermatseev/chat-server/internal/repository/message"
	"github.com/alexandermatseev/chat-server/internal/service"
	chatService "github.com/alexandermatseev/chat-server/internal/service/chat"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig
	authConfig config.AuthServiceConfig

	dbClient              db.Client
	authService           authservice.AuthService
	txManager             db.TxManager
	chatRepository        repository.ChatRepository
	messageRepository     repository.MessageRepository
	contributorRepository repository.ContributorRepository

	chatService service.ChatService

	chatImpl *chat.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) AuthServiceConfig() config.AuthServiceConfig {
	if s.authConfig == nil {
		cfg, err := config.NewAuthServiceConfig()
		if err != nil {
			log.Fatalf("failed to get chat service config: %s", err.Error())
		}

		s.authConfig = cfg
	}

	return s.authConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) AuthService(_ context.Context) authservice.AuthService {
	if s.authService == nil {
		var err error
		s.authService, err = authService2.NewClient(s.AuthServiceConfig().Address())
		if err != nil {
			log.Fatalf("failed to create auth service: %s", err.Error())
		}
	}
	return s.authService
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) MessageRepository(ctx context.Context) repository.MessageRepository {
	if s.messageRepository == nil {
		s.messageRepository = messageRepository.NewRepository(s.DBClient(ctx))
	}

	return s.messageRepository
}

func (s *serviceProvider) ContributorRepository(ctx context.Context) repository.ContributorRepository {
	if s.contributorRepository == nil {
		s.contributorRepository = ContributorRepository.NewRepository(s.DBClient(ctx))
	}

	return s.contributorRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(
			s.ChatRepository(ctx),
			s.MessageRepository(ctx),
			s.ContributorRepository(ctx),
			s.AuthService(ctx),
			s.TxManager(ctx),
		)
	}

	return s.chatService
}

func (s *serviceProvider) ChatImpl(ctx context.Context) *chat.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewImplementation(s.ChatService(ctx))
	}

	return s.chatImpl
}
