package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"tokowiwin/config"
)

type DatabaseRepository struct {
	pgClient *pgxpool.Pool
}

func NewDatabaseRepository(ctx context.Context, cfg *config.AppConfig) *DatabaseRepository {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	if cfg == nil {
		panic("config is nil")
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic occured", r)
		}
	}()
	username := cfg.Database.User
	password := cfg.Database.Pass
	host := cfg.Database.Host
	dbname := cfg.Database.Name
	url := fmt.Sprintf("postgres://%v:%v@%v/%v", username, password, host, dbname)
	conn, err := pgxpool.New(ctx, url)
	if err != nil {
		panic(fmt.Sprintf("error while connect to db client, err=%v", err))
	}
	return &DatabaseRepository{
		pgClient: conn,
	}
}
