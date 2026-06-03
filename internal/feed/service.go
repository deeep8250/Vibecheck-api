package feed

import "github.com/deeep8250/vibecheck-api/internal/models"

type FeedService struct {
	repo FeedRepoInterface
}

func NewFeedService(r FeedRepoInterface) *FeedService {
	return &FeedService{
		repo: r,
	}
}

func (s *FeedService) GetFeed(userID, limit, page int) ([]models.Post, error) {
	offset := (page - 1) * limit
	feed, err := s.repo.GetFeed(userID, limit, offset)
	if err != nil {
		return nil, err
	}
	return feed, nil
}
