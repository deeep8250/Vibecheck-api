package react

import (
	"context"

	"github.com/deeep8250/vibecheck-api/internal/models"
)

type ReactionService struct {
	repo ReactRepositoryInterface
}

func NewReactionService(R ReactRepositoryInterface) *ReactionService {
	return &ReactionService{

		repo: R,
	}
}

func (s *ReactionService) ReactPost(ctx context.Context, r Reaction) error {
	err := s.repo.ReactPost()
}
func GetReactions(ctx context.Context, postID int) ([]models.ReactionDetail, error) {}
