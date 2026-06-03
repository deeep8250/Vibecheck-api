package follow

import "github.com/deeep8250/vibecheck-api/internal/models"

type FollowRepositoryInterface interface {
	FollowRepo(userID, FollowedProfileID int) error
	UnFollowRepo(userID, FollowedProfileID int) error
	GetFollowRepo(userID int) ([]models.Follow, error)
}

type FollowServiceInterface interface {
	FollowService(userID, FollowedProfileID int) error
	UnFollowService(userID, FollowedProfileID int) error
	GetFollowService(userID int) ([]models.Follow, error)
}
