package post

import "github.com/deeep8250/vibecheck-api/internal/models"

type PostRepoInterface interface {
	CreatePost(userID int, userInput CreatePost) (*models.Post, error)
	GetPost(postID int) (*models.Post, error)
	UpdatePost(userID, postID int, userInput UpdatePost) (*models.Post, error)
	DeletePost(userID, postiD int) error
}

type PostServiceInterface interface {
	CreatePostService(userID int, userInput CreatePost) (*models.Post, error)
	GetPostService(postID int) (*models.Post, error)
	UpdatePostService(userID, postID int, userInput UpdatePost) (*models.Post, error)
	DeletePostService(userID, postiD int) error
}
