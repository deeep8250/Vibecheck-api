package react

import (
	"context"

	"github.com/deeep8250/vibecheck-api/internal/models"
)

type ReactServiceInterface interface {
	ReactPost(ctx context.Context, r Reaction) error
	GetReactions(ctx context.Context, postID int) ([]models.ReactionDetail, error)
}

type ReactRepositoryInterface interface {
	ReactPost(ctx context.Context, r Reaction) error
	GetReactions(ctx context.Context, postID int) ([]models.ReactionDetail, error)
	CheckReact(ctx context.Context, postId, userID int) bool
}
