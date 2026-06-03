package post

import (
	"errors"

	"github.com/deeep8250/vibecheck-api/internal/config"
	"github.com/deeep8250/vibecheck-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type PostRepository struct {
	Db *sqlx.DB
}

func NewPostRepository() *PostRepository {
	return &PostRepository{
		Db: config.PostgresDB,
	}
}

func (r *PostRepository) CreatePost(userID int, userInput CreatePost) (*models.Post, error) {

	var Post models.Post
	query := `insert into posts(user_id,content,mood_tag,emoji)  values($1,$2,$3,$4) returning *`
	result := r.Db.QueryRowx(query, userID, userInput.Content, userInput.MoodTag, userInput.Emoji)

	err := result.StructScan(&Post)
	if err != nil {
		return nil, err
	}
	return &Post, nil
}
func (r *PostRepository) GetPost(postID int) (*models.Post, error) {
	var Post models.Post

	query := `select * from posts where id=$1`
	err := r.Db.Get(&Post, query, postID)
	if err != nil {
		return nil, err
	}
	return &Post, nil

}

func (r *PostRepository) GetAllPost(userID int) ([]models.Post, error) {
	var Posts []models.Post

	query := `select * from posts where user_id=$1`
	err := r.Db.Select(&Posts, query, userID)
	if err != nil {
		return nil, err
	}
	return Posts, nil

}

func (r *PostRepository) UpdatePost(userID, postID int, userInput UpdatePost) (*models.Post, error) {
	query := `update posts 
	           set
			      	content = coalesce(nullif($1,''),content),
					mood_tag = coalesce(nullif($2,''),mood_tag),
					emoji = coalesce(nullif($3,''),emoji)
             where id=$4 and user_id=$5 returning *`

	var UpdatedPost models.Post
	result := r.Db.QueryRowx(query, userInput.Content, userInput.MoodTag, userInput.Emoji, postID, userID)
	err := result.StructScan(&UpdatedPost)
	if err != nil {
		return nil, err
	}
	return &UpdatedPost, nil

}
func (r *PostRepository) DeletePost(userID, postiD int) error {

	query := `delete from posts where id=$1 and user_id=$2`
	result, err := r.Db.Exec(query, postiD, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {

		return errors.New("post not found or unauthorized")
	}

	return nil

}
