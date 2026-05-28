package database

import (
	"context"
	"fmt"
	"log"

	"github.com/HR-Shekhar/todo-api/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresConnection(cfg *config.Config) *pgxpool.Pool {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)
	dbpool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		log.Fatal("Unable to create connection pool:\n", err)
	}
	if err := dbpool.Ping(context.Background()); err != nil {
		log.Fatal("Connection to db failed")
	} 

	return dbpool
}
