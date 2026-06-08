package post

import (
	"context"

	"github.com/deeep8250/vibecheck-api/internal/models"
)

type PostRepoInterface interface {
	CreatePost(ctx context.Context, userID int, userInput CreatePost) (*models.Post, error)
	GetAllPost(ctx context.Context, userID int) ([]models.Post, error)
	GetPost(ctx context.Context, postID int) (*models.Post, error)
	UpdatePost(ctx context.Context, userID, postID int, userInput UpdatePost) (*models.Post, error)
	DeletePost(ctx context.Context, userID, postiD int) error
}

type PostServiceInterface interface {
	CreatePostService(ctx context.Context, userID int, userInput CreatePost) (*models.Post, error)
	GetPostService(ctx context.Context, postID int) (*models.Post, error)
	GetAllPostService(ctx context.Context, userID int) ([]models.Post, error)
	UpdatePostService(ctx context.Context, userID, postID int, userInput UpdatePost) (*models.Post, error)
	DeletePostService(ctx context.Context, userID, postiD int) error
}
