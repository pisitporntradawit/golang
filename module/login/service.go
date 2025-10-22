package login

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("Secret")

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service{
	return &Service{
		repo: repo,
	}
}

func(s *Service) GetLogin(name,password string) (string,error){
	user, err := s.repo.GetLogin(name)
	if err != nil {
		return "",errors.New("User not found")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"id" : user.ID,
		"exp" : time.Now().Add(time.Hour*24).Unix(),
	})

	tokenString,err := token.SignedString(jwtSecret)
	if err != nil{
		return "",err
	}

	return tokenString,nil
}