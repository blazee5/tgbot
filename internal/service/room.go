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
	repo repository.RoomRepository
}

func NewRoomService(log *zap.SugaredLogger, repo repository.RoomRepository) *RoomService {
	return &RoomService{log: log, repo: repo}
}

func (s *RoomService) BookRoom(ctx context.Context, input models.Book) error {
	ok, err := s.repo.GetBook(ctx, input)

	if err != nil {
		s.log.Infof("error while get book: %v", err)
	}

	if ok {
		return bot_errors.RoomAlreadyBookedErr
	}

	return s.repo.SetBook(ctx, input, 3600)
}
