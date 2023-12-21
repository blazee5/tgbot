package repository

import (
	"context"
	"github.com/blazee5/tgbot/internal/models"
	"github.com/blazee5/tgbot/internal/repository/postgres"
	redisRepo "github.com/blazee5/tgbot/internal/repository/redis"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	UserRepository
	RoomRepository
}

func NewRepository(db *pgxpool.Pool, rdb *redis.Client) *Repository {
	return &Repository{
		UserRepository: postgres.NewUserPostgres(db),
		RoomRepository: redisRepo.NewRoomRepository(rdb),
	}
}

type UserRepository interface {
	Create(ctx context.Context, input models.User) (int, error)
	GetByID(ctx context.Context, id int) (models.User, error)
	VerifyUser(ctx context.Context, id int) error
}

type RoomRepository interface {
	SetBook(ctx context.Context, input models.Book, seconds int) error
	GetBook(ctx context.Context, input models.Book) (bool, error)
}
