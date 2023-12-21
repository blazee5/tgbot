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

func (s *UserService) CreateUser(ctx context.Context, input models.User) (int, error) {
	user, err := s.repo.GetByID(ctx, input.ID)

	if err != nil {
	}

	var id int

	if user.ID == 0 {
		id, err = s.repo.Create(ctx, input)

		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (s *UserService) GetByID(ctx context.Context, id int) (models.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) VerifyUser(ctx context.Context, id int) error {
	return s.repo.VerifyUser(ctx, id)
}
