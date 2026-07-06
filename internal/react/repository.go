package react

import (
	"context"
	"errors"

	"github.com/deeep8250/vibecheck-api/internal/config"
	"github.com/deeep8250/vibecheck-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type ReactRepository struct {
	db *sqlx.DB
}

func NewReactRepository() *ReactRepository {
	return &ReactRepository{
		db: config.PostgresDB,
	}
}

func (r *ReactRepository) ReactPost(ctx context.Context, postID, userID int, emoji string) error {

	query := `insert into reactions(post_id,reaction_giver_id,reaction_emoji) values($1,$2,$3)`
	_, err := r.db.ExecContext(ctx, query, postID, userID, emoji)
	if err != nil {
		return err
	}

	return nil

}
func (r *ReactRepository) GetReactionsByPostID(ctx context.Context, postID int) ([]ReactionDetail, error) {
	query := `SELECT u.username,r.id,r.reaction_emoji,r.created_at from users as u
	join reactions as r on r.reaction_giver_id=u.id where r.post_id=$1
   
	`

	var PostsWithReact []ReactionDetail

	err := r.db.SelectContext(ctx, &PostsWithReact, query, postID)
	if err != nil {
		return nil, err
	}
	return PostsWithReact, nil

}

func (r *ReactRepository) GetPost(ctx context.Context, postID int) (*models.Post, error) {
	query := `select * from posts where id=$1`

	var post models.Post
	err := r.db.GetContext(ctx, &post, query, postID)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *ReactRepository) DeleteReact(ctx context.Context, postId, userId int) error {
	query := `delete from reactions where reaction_giver_id=$1 and post_id=$2`
	result, err := r.db.ExecContext(ctx, query, userId, postId)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("reaction not found or unauthorized")
	}
	return nil
}
