package auth

import (
	"github.com/deeep8250/vibecheck-api/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo AuthRepoInterface
}

func NewAuthService(Repo AuthRepoInterface) *AuthService {
	return &AuthService{
		repo: Repo,
	}
}

func (s *AuthService) Register(userInput Register) (*models.User, error) {

	hashPass, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	userInput.Password = string(hashPass)

	user, err := s.repo.CreateUser(userInput.Username, userInput.Email, userInput.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(userInput Login) (string, error) {
	user, err := s.repo.GetUserByEmail(userInput.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(userInput.Password))
	if err != nil {
		return "", err
	}

	token, err := JWTinit(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
