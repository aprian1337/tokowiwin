package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"strings"
	"tokowiwin/constants"
	"tokowiwin/repositories/db/queries"
	"tokowiwin/repositories/model"
	"tokowiwin/utils/array"
	"tokowiwin/utils/converts"
	"tokowiwin/utils/structs"
)

type PsqlRepo interface {
	GetUserByEmail(ctx context.Context, email string, columns ...any) (*model.Users, error)
	InsertEmail(ctx context.Context, tx pgx.Tx, users *model.Users) error
}

func (r *DatabaseRepository) GetUserByEmail(ctx context.Context, email string, columns ...any) (*model.Users, error) {
	var (
		err     error
		modelDb = new(model.Users)
	)

	if len(columns) == 0 {
		columns = structs.GetColumns(&model.Users{})
	}

	columnsStr := array.ToStringWithDelimiter(columns, ",")
	query := strings.ReplaceAll(queries.QueryUsersByEmail, constants.DB_COLS, columnsStr)

	err = r.pgClient.QueryRow(ctx, query, email).Scan(structs.GetAddressByFieldTag(modelDb, constants.DB_TAG, converts.AnyArrayToString(columns))...)
	if err != nil {
		return nil, err
	}

	return modelDb, nil
}

func (r *DatabaseRepository) InsertEmail(ctx context.Context, tx pgx.Tx, users *model.Users) error {
	var (
		err error
	)

	if tx == nil {
		tx, err = r.pgClient.Begin(ctx)
		if err != nil {
			return err
		}
	}
	columns := structs.GetColumns(&model.Users{})
	columnsStr := array.ToStringWithDelimiter(columns, ",")
	query := strings.ReplaceAll(queries.QueryInsertUsers, constants.DB_COLS, columnsStr)
	fmt.Println("Q", query)
	_, err = r.pgClient.Exec(ctx, query, structs.GetAddressByFieldTag(users, constants.DB_TAG, converts.AnyArrayToString(columns))...)
	if err != nil {
		return err
	}

	return nil
}
