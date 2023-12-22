package redis

import (
	"context"
	"fmt"
	"github.com/blazee5/tgbot/internal/models"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

type RoomRepository struct {
	client *redis.Client
}

func NewRoomRepository(client *redis.Client) *RoomRepository {
	return &RoomRepository{client: client}
}

func (repo *RoomRepository) SetBook(ctx context.Context, input models.Book, seconds int) error {
	key := fmt.Sprintf("%s:%s:%d", input.RoomId, input.Time, input.UserId)

	err := repo.client.Set(ctx, key, input.Username, time.Duration(seconds)*time.Second).Err()

	if err != nil {
		return err
	}

	return nil
}

func (repo *RoomRepository) GetBook(ctx context.Context, room, hour string) (string, error) {
	var username string

	key := fmt.Sprintf("%s:%s*", room, hour)

	keys, err := repo.client.Keys(ctx, key).Result()

	if err != nil {
		return "", err
	}

	for _, k := range keys {
		err = repo.client.Get(ctx, k).Scan(&username)

		if err != nil {
			return "", err
		}
	}

	return username, nil
}

func (repo *RoomRepository) GetBooksByID(ctx context.Context, userID int) ([]models.Book, error) {
	books := make([]models.Book, 0)

	key := fmt.Sprintf("*%d", userID)

	keys, err := repo.client.Keys(ctx, key).Result()

	if err != nil {
		return nil, err
	}

	for _, k := range keys {
		parts := strings.Split(k, ":")

		books = append(books, models.Book{
			RoomId: parts[0],
			Time:   parts[1],
		})
	}

	return books, nil
}

func (repo *RoomRepository) DeleteBook(ctx context.Context, roomId, hour string, userID int) error {
	key := fmt.Sprintf("%s:%s:%d", roomId, hour, userID)

	err := repo.client.Del(ctx, key).Err()

	if err != nil {
		return err
	}

	return nil
}
