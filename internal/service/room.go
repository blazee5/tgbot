package service

import (
	"context"
	"github.com/blazee5/tgbot/internal/models"
	"github.com/blazee5/tgbot/internal/repository"
	"github.com/blazee5/tgbot/pkg/bot_errors"
	"go.uber.org/zap"
)

type RoomService struct {
	log  *zap.SugaredLogger
	repo *repository.Repository
}

func NewRoomService(log *zap.SugaredLogger, repo *repository.Repository) *RoomService {
	return &RoomService{log: log, repo: repo}
}

func (s *RoomService) GetBooksByID(ctx context.Context, userID int) ([]models.Book, error) {
	return s.repo.RoomRepository.GetBooksByID(ctx, userID)
}

func (s *RoomService) BookRoom(ctx context.Context, userID int, room, hour string) (string, error) {
	username, err := s.repo.RoomRepository.GetBook(ctx, room, hour)

	if err != nil {
		s.log.Infof("error while get book: %v", err)
	}

	if username != "" {
		return username, bot_errors.RoomAlreadyBookedErr
	}

	user, err := s.repo.UserRepository.GetByID(ctx, userID)

	if err != nil {
		return "", err
	}

	return "", s.repo.SetBook(ctx, models.Book{
		UserId:   userID,
		Username: user.FirstName,
		RoomId:   room,
		Time:     hour,
	}, 3600)
}

func (s *RoomService) CancelBook(ctx context.Context, roomID, hour string, userID int) error {
	return s.repo.RoomRepository.DeleteBook(ctx, roomID, hour, userID)
}
