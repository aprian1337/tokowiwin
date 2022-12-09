package db

import (
	"context"
	"strings"
	"tokowiwin/constants"
	"tokowiwin/repositories/db/queries"
	"tokowiwin/repositories/model"
	"tokowiwin/utils/array"
	"tokowiwin/utils/converts"
	"tokowiwin/utils/structs"
)

type PsqlRepo interface {
	GetUserByEmail(ctx context.Context, email string, columns ...string) (error, *model.Users)
}

func (r *DatabaseRepository) GetUserByEmail(ctx context.Context, email string, columns ...any) (*model.Users, error) {
	var modelDb = new(model.Users)

	if len(columns) == 0 {
		columns = structs.GetColumns(&model.Users{})
	}

	columnsStr := array.ToStringWithDelimiter(columns, ",")
	query := strings.ReplaceAll(queries.QueryUsersByEmail, constants.DB_COLS, columnsStr)

	err := r.pgClient.QueryRow(ctx, query, email).Scan(structs.GetAddressByFieldTag(modelDb, constants.DB_TAG, converts.AnyArrayToString(columns))...)
	if err != nil {
		return nil, err
	}

	return modelDb, nil
}
