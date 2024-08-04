package converter

import (
	"github.com/alexandermatseev/chat-server/internal/model"
)

// ToContributorsCreateFromcht converts cht.ContributorsCreate to model.ContributorsCreate
func ToContributorsCreateFromcht(userIDs []int64) *model.ContributorsCreate {
	ContributorsCreate := &model.ContributorsCreate{
		Contributors: make([]model.ContributorCreate, 0, len(userIDs)),
	}

	for _, userID := range userIDs {
		ContributorsCreate.Contributors = append(ContributorsCreate.Contributors, model.ContributorCreate{
			UserID: userID,
		})
	}

	return ContributorsCreate
}
