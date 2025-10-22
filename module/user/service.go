package user

import (
	"context"
	"fmt"
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

func (s *Service) GetUser(ctx context.Context) ([]GetUserModel, error) {
	return s.repo.GetUser(ctx)
}

func (s *Service) GetUserByID(ctx context.Context, id string) (*GetUserModel, error){
	return s.repo.GetUserByID(ctx, id)
}

func (s *Service) InsertUser(ctx context.Context, user *User, profile *Profile) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error Hash ")
	}
	user.Password = string(hashPassword)
	err = s.repo.InsertUser(ctx, user, profile)
	if err != nil {
		fmt.Println("InsertUser failed: %w", err)
	}
	return nil
}

func(s *Service) DeleteUser(ctx context.Context, id string) error {
	return s.repo.DeleteUser(ctx, id)
}
