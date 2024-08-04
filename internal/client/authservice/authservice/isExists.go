package authservice

import (
	"context"
)

// IsUserExists реализует метод интерфейса AuthService
func (c *Client) IsUserExists(ctx context.Context, userIDs []int64) (bool, error) {
	return true, nil
}
