package feed

import (
	"context"

	"github.com/deeep8250/vibecheck-api/internal/models"
)

type FeedRepoInterface interface {
	GetFeed(ctx context.Context, userID, limit, offset int) ([]models.Post, error)
}

type FeedServiceInterface interface {
	GetFeed(ctx context.Context, userID, limit, page int) ([]models.Post, error)
}
