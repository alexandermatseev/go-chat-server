package model

// Contributor represents a chat Contributor
type Contributor struct {
	ChatID int64
	UserID int64
}

// ContributorCreate represents a chat Contributor to be created
type ContributorCreate struct {
	ChatID int64
	UserID int64
}

// ContributorsCreate represents a chat Contributors to be created
type ContributorsCreate struct {
	Contributors []ContributorCreate
}
