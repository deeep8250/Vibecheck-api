package feed

import (
	"github.com/deeep8250/vibecheck-api/internal/config"
	"github.com/deeep8250/vibecheck-api/internal/models"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type FeedRepo struct {
	db *sqlx.DB
}

func NewFeedRepo() *FeedRepo {
	return &FeedRepo{
		db: config.PostgresDB,
	}
}

func (r *FeedRepo) GetFeed(c context.Context, userID, limit, offset int) ([]models.Post, error) {
	var Feed []models.Post

	query := `select p.*,count(r.id) as reaction_count from posts as p
	join follows as f on f.followed_user_id=p.user_id
	left join reactions as r on p.id=r.post_id
	where f.follower_id=$1 
	group by p.id
	order by p.created_at desc
	limit $2 offset $3`

	err := r.db.SelectContext(c, &Feed, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	return Feed, nil

}
