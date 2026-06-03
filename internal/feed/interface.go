package feed

import "github.com/deeep8250/vibecheck-api/internal/models"

type FeedRepoInterface interface {
	GetFeed(userID, limit, offset int) ([]models.Post, error)
}

type FeedServiceInterface interface {
	GetFeed(userID, limit, page int) ([]models.Post, error)
}
