package feed

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/deeep8250/vibecheck-api/internal/config"
	"github.com/deeep8250/vibecheck-api/internal/models"
	"github.com/redis/go-redis/v9"
)

type FeedService struct {
	repo  FeedRepoInterface
	redis *redis.Client
}

func NewFeedService(r FeedRepoInterface) *FeedService {
	return &FeedService{
		repo:  r,
		redis: config.RedisClient,
	}
}

func (s *FeedService) GetFeed(c context.Context, userID, limit, page int) ([]models.Post, error) {

	// build cache key
	cacheKey := fmt.Sprintf("feed:%d:%d:%d", userID, page, limit)

	//checking redis
	cache, err := s.redis.Get(c, cacheKey).Result()
	if err == nil {
		var feed []models.Post
		if err := json.Unmarshal([]byte(cache), &feed); err == nil {
			return feed, nil
		}
	}

	// 3. cache miss — fetch from DB
	offset := (page - 1) * limit
	feed, err := s.repo.GetFeed(c, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	// 4. store in Redis with TTL
	data, err := json.Marshal(feed)
	if err != nil {
		return nil, err
	}
	s.redis.Set(c, cacheKey, data, time.Minute)

	return feed, nil
}
