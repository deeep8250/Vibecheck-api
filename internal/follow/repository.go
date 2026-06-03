package follow

import (
	"errors"

	"github.com/deeep8250/vibecheck-api/internal/config"
	"github.com/deeep8250/vibecheck-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type FollowRepository struct {
	Db *sqlx.DB
}

func NewFollowRepository() *FollowRepository {
	return &FollowRepository{
		Db: config.PostgresDB,
	}
}

func (r *FollowRepository) FollowRepo(userID, FollowedProfileID int) error {

	query := `insert into follows(follower_id,followed_user_id) values($1,$2)`
	result, err := r.Db.Exec(query, userID, FollowedProfileID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("cant follow the user something went wrong")
	}

	return nil

}

func (r *FollowRepository) UnFollowRepo(userID, FollowedProfileID int) error {

	query := `delete from follows where follower_id=$1 and followed_user_id=$2`
	result, err := r.Db.Exec(query, userID, FollowedProfileID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("failed to unfollow something went wrong")
	}

	return nil

}
func (r *FollowRepository) GetFollowRepo(userID int) ([]models.Follow, error) {
	var followedProfiles []models.Follow
	query := `select * from follows where follower_id=$1`
	err := r.Db.Select(&followedProfiles, query, userID)
	if err != nil {
		return nil, err
	}

	return followedProfiles, nil
}
