package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type DBConfig struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

func NewPostgresDB(cfg DBConfig) (*pgx.Conn, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Name)
	db, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return db, nil
}
