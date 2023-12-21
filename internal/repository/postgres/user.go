package postgres

import (
	"context"
	"github.com/blazee5/tgbot/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserPostgres struct {
	db *pgxpool.Pool
}

func NewUserPostgres(db *pgxpool.Pool) *UserPostgres {
	return &UserPostgres{db: db}
}

func (repo *UserPostgres) Create(ctx context.Context, input models.User) (int, error) {
	var id int

	err := repo.db.QueryRow(ctx, "INSERT INTO users (id, first_name, last_name, username) VALUES ($1, $2, $3, $4) RETURNING id",
		input.ID, input.FirstName, input.LastName, input.Username).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (repo *UserPostgres) GetByID(ctx context.Context, id int) (models.User, error) {
	var user models.User

	err := repo.db.QueryRow(ctx, "SELECT id, first_name, last_name, username, is_verified FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.IsVerified)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (repo *UserPostgres) VerifyUser(ctx context.Context, id int) error {
	_, err := repo.db.Exec(ctx, "UPDATE users SET is_verified = true WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}
