package auth

import (
	"context"

	"github.com/deeep8250/vibecheck-api/internal/models"
)

// for handler test
type AuthServiceInterface interface {
	Register(c context.Context, r Register) (*models.User, error)
	Login(c context.Context, userInput Login) (string, error)
}

// for sevice test
type AuthRepoInterface interface {
	CreateUser(c context.Context, username, email, password string) (*models.User, error)
	GetUserByEmail(c context.Context, email string) (*models.User, error)
}
