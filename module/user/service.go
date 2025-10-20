package user

import (
	"context"
	"fmt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetUser(ctx context.Context) ([]User, error) {
	return s.repo.GetUser(ctx)
}

func (s *Service) InsertUser(ctx context.Context, user *User) error {
	err := s.repo.InsertUser(ctx, user)
	if err != nil {
		fmt.Println("InsertUser failed: %w", err)
	}
	return nil
}

func(s *Service) DeleteUser(ctx context.Context, id string) error {
	return s.repo.DeleteUser(ctx, id)
}
