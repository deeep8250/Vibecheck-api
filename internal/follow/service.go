package follow

import (
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

func (s *FollowService) FollowService(userID, FollowedProfileID int) error {
	if userID == FollowedProfileID {
		return errors.New("cant follow yourself dumbass")
	}
	err := s.repo.FollowRepo(userID, FollowedProfileID)
	if err != nil {
		return err
	}
	return nil

}
func (s *FollowService) UnFollowService(userID, FollowedProfileID int) error {
	err := s.repo.UnFollowRepo(userID, FollowedProfileID)
	if err != nil {
		return err
	}
	return nil

}
func (s *FollowService) GetFollowService(userID int) ([]models.Follow, error) {
	following, err := s.repo.GetFollowRepo(userID)
	if err != nil {
		return nil, err
	}
	return following, nil
}
