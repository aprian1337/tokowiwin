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

var client *pgxpool.Pool

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
	dbConfig := cfg.Database
	url := fmt.Sprintf("postgres://%v:%v@%v/%v", dbConfig.User, dbConfig.Pass, dbConfig.Host, dbConfig.Name)
	conn, err := pgxpool.New(ctx, url)
	if err != nil {
		panic(fmt.Sprintf("error while connect to db client, err=%v", err))
	}
	_, err = conn.Acquire(ctx)
	if err != nil {
		return nil
	}
	client = conn
	return &DatabaseRepository{
		pgClient: conn,
	}
}

func GetDBClient() *pgxpool.Pool {
	return client
}
