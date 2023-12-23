package service

import (
	"context"
	"github.com/blazee5/tgbot/internal/models"
	"github.com/blazee5/tgbot/internal/repository"
	"go.uber.org/zap"
)

type Service struct {
	User
	Room
}

type User interface {
	CreateUser(ctx context.Context, input models.User) error
	GetByID(ctx context.Context, id int) (models.User, error)
	VerifyUser(ctx context.Context, id int) error
}

type Room interface {
	BookRoom(ctx context.Context, userID int, room, hour string) (string, error)
	GetBook(ctx context.Context, room, hour string) (string, error)
	GetBooksByID(ctx context.Context, userID int) ([]models.Book, error)
	CancelBook(ctx context.Context, roomID, hour string, userID int) error
}

func NewService(log *zap.SugaredLogger, repo *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repo),
		Room: NewRoomService(log, repo),
	}
}
