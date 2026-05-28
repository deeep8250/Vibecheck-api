package auth

import (
	"github.com/deeep8250/vibecheck-api/internal/config"
	"github.com/deeep8250/vibecheck-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	DbClient *sqlx.DB
}

func NewDb() *AuthRepository {
	return &AuthRepository{
		DbClient: config.PostgresDB,
	}
}

func (db *AuthRepository) CreateUser(username, email, passwordHash string) (*models.User, error) {
	var User models.User
	// returning *  return the inserted row
	query := `insert into users(username, email, password_hash) values($1,$2,$3) returning *`

	// exec not return the row inserted thats why we use queryrowx to get the inserted row but we need to use struct scan with that to catch the return item
	result := db.DbClient.QueryRowx(query, username, email, passwordHash)
	err := result.StructScan(&User)
	if err != nil {
		return nil, err
	}
	return &User, nil
}

func (db *AuthRepository) GetUserByEmail(email string) (*models.User, error) {
	var User models.User
	query := `select * from users where email=$1`

	result := db.DbClient.QueryRowx(query, email)
	err := result.StructScan(&User)
	if err != nil {
		return nil, err
	}
	return &User, nil
}
