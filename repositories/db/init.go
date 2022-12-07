package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
	"tokowiwin/constants"
)

type DatabaseRepository struct {
	PgClient *pgxpool.Pool
}

func NewDatabaseRepository(ctx context.Context) *DatabaseRepository {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic occured", r)
		}
	}()
	username := viper.GetString(constants.DATABASE_USER)
	password := viper.GetString(constants.DATABASE_PASS)
	host := viper.GetString(constants.DATABASE_HOST)
	dbname := viper.GetString(constants.DATABASE_NAME)
	url := fmt.Sprintf("postgres://%v:%v@%v/%v", username, password, host, dbname)
	conn, err := pgxpool.New(ctx, url)
	if err != nil {
		panic(fmt.Sprintf("error while connect to db client, err=%v", err))
	}
	return &DatabaseRepository{
		PgClient: conn,
	}
}
