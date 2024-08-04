package chat

import (
	"context"
	"errors"
	"github.com/alexandermatseev/chat-server/internal/model"
)

func (s *serv) Create(ctx context.Context, createChat *model.ChatCreate, createContributors *model.ContributorsCreate) (int64, error) {
	var id int64

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		userIDs := make([]int64, 0, len(createContributors.Contributors))
		for _, Contributor := range createContributors.Contributors {
			userIDs = append(userIDs, Contributor.UserID)
		}
		exists, errTx := s.authServiceClient.IsUserExists(ctx, userIDs)
		if errTx != nil {
			return errTx
		}
		if !exists {
			return errors.New("some users do not exist")
		}

		id, errTx = s.chatRepository.Create(ctx, createChat)
		if errTx != nil {
			return errTx
		}
		for i := 0; i < len(createContributors.Contributors); i++ {
			createContributors.Contributors[i].ChatID = id
		}

		errTx = s.ContributorRepository.CreateContributors(ctx, createContributors)
		if errTx != nil {
			return errTx
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}
