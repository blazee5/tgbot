package redis

import (
	"context"
	"fmt"
	"github.com/blazee5/tgbot/internal/models"
	"github.com/redis/go-redis/v9"
	"time"
)

type RoomRepository struct {
	client *redis.Client
}

func NewRoomRepository(client *redis.Client) *RoomRepository {
	return &RoomRepository{client: client}
}

func (r *RoomRepository) SetBook(ctx context.Context, input models.Book, seconds int) error {
	key := fmt.Sprintf("%d:%s:%s", input.UserID, input.RoomId, input.Time)

	err := r.client.Set(ctx, key, true, time.Duration(seconds)*time.Second).Err()

	if err != nil {
		return err
	}

	return nil
}

func (r *RoomRepository) GetBook(ctx context.Context, input models.Book) (bool, error) {
	var ok bool

	key := fmt.Sprintf("%d:%s:%s", input.UserID, input.RoomId, input.Time)

	err := r.client.Get(ctx, key).Scan(&ok)

	if err != nil {
		return false, err
	}

	return ok, nil
}
