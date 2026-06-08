package follow

import (
	"context"

	"github.com/deeep8250/vibecheck-api/internal/models"
)

type FollowRepositoryInterface interface {
	FollowRepo(c context.Context, userID, FollowedProfileID int) error
	UnFollowRepo(c context.Context, userID, FollowedProfileID int) error
	GetFollowRepo(c context.Context, userID int) ([]models.Follow, error)
}

type FollowServiceInterface interface {
	FollowService(c context.Context, userID, FollowedProfileID int) error
	UnFollowService(c context.Context, userID, FollowedProfileID int) error
	GetFollowService(c context.Context, userID int) ([]models.Follow, error)
}
