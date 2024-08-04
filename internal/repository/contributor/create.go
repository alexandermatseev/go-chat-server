package contributor

import (
	"context"

	"github.com/alexandermatseev/platform_common/pkg/db"

	sq "github.com/Masterminds/squirrel"
	"github.com/alexandermatseev/chat-server/internal/model"
)

func (r *repo) CreateContributor(ctx context.Context, createContributor *model.ContributorCreate) error {
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIDColumn, userIDColumn).
		Values(createContributor.ChatID, createContributor.UserID)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "Contributor_repository.CreateContributor",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) CreateContributors(ctx context.Context, createContributors *model.ContributorsCreate) error {
	for i := 0; i < len(createContributors.Contributors); i++ {
		err := r.CreateContributor(ctx, &createContributors.Contributors[i])
		if err != nil {
			return err
		}
	}
	return nil
}
