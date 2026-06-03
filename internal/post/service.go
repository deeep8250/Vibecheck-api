package post

import "github.com/deeep8250/vibecheck-api/internal/models"

type PostService struct {
	repo PostRepoInterface
}

func NewPostService(R PostRepoInterface) *PostService {
	return &PostService{
		repo: R,
	}
}

func (r *PostService) CreatePostService(userID int, userInput CreatePost) (*models.Post, error) {
	user, err := r.repo.CreatePost(userID, userInput)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PostService) GetPostService(postID int) (*models.Post, error) {
	user, err := r.repo.GetPost(postID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *PostService) GetAllPostService(userID int) ([]models.Post, error) {
	userPosts, err := r.repo.GetAllPost(userID)
	if err != nil {
		return nil, err
	}
	return userPosts, nil
}

func (r *PostService) UpdatePostService(userID, postID int, userInput UpdatePost) (*models.Post, error) {
	user, err := r.repo.UpdatePost(userID, postID, userInput)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *PostService) DeletePostService(userID, postiD int) error {
	err := r.repo.DeletePost(userID, postiD)
	if err != nil {
		return err
	}
	return nil
}
