package post

import "github.com/deeep8250/vibecheck-api/internal/models"

type PostRepoInterface interface {
	CreatePost(userInput CreatePost) (*models.Post, error)
	GetPost(postID int) (*models.Post, error)
	UpdatePost(userInput UpdatePost) (*models.Post, error)
	DeletePost(userID, postiD int) error
}

type PostServiceInterface interface {
	CreatePost(userInput CreatePost) (*models.Post, error)
	GetPost(postID int) (*models.Post, error)
	UpdatePost(userInput UpdatePost) (*models.Post, error)
	DeletePost(userID, postiD int) error
}
