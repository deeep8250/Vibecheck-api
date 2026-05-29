package auth

import (
	"github.com/deeep8250/vibecheck-api/internal/models"
)

// for handler test
type AuthServiceInterface interface {
	Register(Register) (*models.User, error)
	Login(userInput Login) (string, error)
}

// for sevice test
type AuthRepoInterface interface {
	CreateUser(username, email, password string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}
