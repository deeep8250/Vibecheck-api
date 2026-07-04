package react

import (
	"context"

	"github.com/deeep8250/vibecheck-api/internal/models"
)

type ReactRepoInterface interface {
	ReactPost(ctx context.Context, postID, userID int, emoji string) error
	GetReactionsByPostID(ctx context.Context, postID int) ([]ReactionDetail, error)
	GetPost(ctx context.Context, postID int) (*models.Post, error)
	DeleteReact(ctx context.Context, postId, userId int) error
}

type ReactServiceInterface interface {
	ReactPostServ(ctx context.Context, postID, userID int, emoji string) error
	GetReactionsByPostID(ctx context.Context, postID int) (*PostWithReactions, error)
	RemoveReact(ctx context.Context, postID, userID int) error
}
