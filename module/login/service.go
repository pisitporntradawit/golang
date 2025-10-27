package login

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetLogin(ctx context.Context, username, password string) (string, error) {

	user, err := s.repo.GetLogin(ctx, username)
	if err != nil {
		return "", errors.New("user not found")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", errors.New("invalid password")
	}
	role, err := s.repo.AuthRole(ctx, user.ID.String())
	if err != nil {
		return "", errors.New("role not found")
	}
	//สร้าง Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       role.ID,
		"username": role.Username,
		"role":     role.Rolename,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWTSECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
