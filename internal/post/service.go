package post

import (
	"context"

	"github.com/deeep8250/vibecheck-api/internal/models"
)

type PostService struct {
	repo PostRepoInterface
}

func NewPostService(R PostRepoInterface) *PostService {
	return &PostService{
		repo: R,
	}
}

func (r *PostService) CreatePostService(ctx context.Context, userID int, userInput CreatePost) (*models.Post, error) {
	user, err := r.repo.CreatePost(ctx, userID, userInput)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PostService) GetPostService(ctx context.Context, postID int) (*models.Post, error) {
	user, err := r.repo.GetPost(ctx, postID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *PostService) GetAllPostService(ctx context.Context, userID int) ([]models.Post, error) {
	userPosts, err := r.repo.GetAllPost(ctx, userID)
	if err != nil {
		return nil, err
	}
	return userPosts, nil
}

func (r *PostService) UpdatePostService(ctx context.Context, userID, postID int, userInput UpdatePost) (*models.Post, error) {
	user, err := r.repo.UpdatePost(ctx, userID, postID, userInput)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *PostService) DeletePostService(ctx context.Context, userID, postiD int) error {
	err := r.repo.DeletePost(ctx, userID, postiD)
	if err != nil {
		return err
	}
	return nil
}
