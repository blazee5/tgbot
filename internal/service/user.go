package service

import (
	"context"
	"github.com/blazee5/tgbot/internal/models"
	"github.com/blazee5/tgbot/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, input models.User) error {
	_, err := s.repo.Create(ctx, input)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetByID(ctx context.Context, id int) (models.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) VerifyUser(ctx context.Context, id int) error {
	return s.repo.VerifyUser(ctx, id)
}
