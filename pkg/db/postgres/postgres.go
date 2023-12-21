package postgres

import (
	"context"
	"fmt"
	"github.com/blazee5/tgbot/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func NewPostgres(cfg *config.Config) *pgxpool.Pool {
	db, err := pgxpool.New(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name))

	if err != nil {
		log.Fatalf("error while connect to postgres: %v", err)
	}

	if err = db.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	return db
}
