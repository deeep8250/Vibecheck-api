package react

import (
	"errors"

	"github.com/deeep8250/vibecheck-api/internal/config"
	"github.com/deeep8250/vibecheck-api/internal/models"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type ReactRepo struct {
	db *sqlx.DB
}

func NewReactRepo() *ReactRepo {
	return &ReactRepo{
		db: config.PostgresDB,
	}
}

func (r *ReactRepo) AddReaction(ctx context.Context, input Reaction) error {
	query := `insert into reaction (post_id,reaction_giver_id) values ($1,$2)`
	rowsAffected, err := r.db.ExecContext(ctx, query, input.PostID, input.Reaction)
	if err != nil {
		return err
	}
	RowsAff, _ := rowsAffected.RowsAffected()
	if RowsAff == 0 {
		return errors.New("unable to add your reaction")
	}
	return nil
}

func (r *ReactRepo) GetReactions(ctx context.Context, postID int) ([]models.ReactionDetail, error) {
	query := `select u.name ,p.content,p.mood_tag,p.emoji,p.post_date,p.created_at,r.reaction_emoji  from users as u 
	          join posts as p  on p.user_id=u.id
			  join reactions as r on r.post_id=p.id
			  where r.post_id=$1 
			  GROUP BY 
			u.name,
			p.content,
			p.mood_tag,
			p.emoji,
			p.post_date,
			p.created_at,
			r.reaction_emoji
			  `

	var Reactions []models.ReactionDetail
	err := r.db.SelectContext(ctx, &Reactions, query, postID)
	if err != nil {
		return nil, err
	}
	return Reactions, nil
}
func (r *ReactRepo) CheckReact(ctx context.Context, postId, userID int) (int, error) {

	query := `select count(*) from reactions where post_id=$1 and  reaction_giver_id=$2 `

	var cnt int
	err := r.db.GetContext(ctx, &cnt, query, postId, userID)

	if err != nil {
		return 0, err
	}
	if cnt <= 0 {
		return 0, errors.New("no reaction found in this post")
	}
	return cnt, nil

}
