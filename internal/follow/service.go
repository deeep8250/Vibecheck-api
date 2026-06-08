package follow

import (
	"context"
	"errors"

	"github.com/deeep8250/vibecheck-api/internal/models"
)

type FollowService struct {
	repo FollowRepositoryInterface
}

func NewFollowService(re FollowRepositoryInterface) *FollowService {
	return &FollowService{
		repo: re,
	}
}

func (s *FollowService) FollowService(ctx context.Context, userID, FollowedProfileID int) error {

	if userID == FollowedProfileID {
		return errors.New("cant follow yourself dumbass")
	}
	err := s.repo.FollowRepo(ctx, userID, FollowedProfileID)
	if err != nil {
		return err
	}
	return nil

}
func (s *FollowService) UnFollowService(ctx context.Context, userID, FollowedProfileID int) error {
	err := s.repo.UnFollowRepo(ctx, userID, FollowedProfileID)
	if err != nil {
		return err
	}
	return nil

}
func (s *FollowService) GetFollowService(ctx context.Context, userID int) ([]models.Follow, error) {
	following, err := s.repo.GetFollowRepo(ctx, userID)
	if err != nil {
		return nil, err
	}
	return following, nil
}
