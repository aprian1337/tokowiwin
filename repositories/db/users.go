package db

import (
	"context"
	"strings"
	"tokowiwin/constants"
	"tokowiwin/repositories/model"
	"tokowiwin/utils/array"
	"tokowiwin/utils/converts"
	"tokowiwin/utils/structs"
)

type AuthenticationRepo interface {
	GetUserByEmail(ctx context.Context, email string, columns ...string) (error, *model.Users)
}

const (
	queryLogin = "SELECT ${cols} FROM users WHERE email=$1"
)

//func (r *DatabaseRepository) GetUserByEmail(ctx context.Context, email string, columns ...any) ([]*model.Users, error) {
//	var result []*model.Users
//
//	if len(columns) == 0 {
//		columns = structs.GetColumns(&model.Users{})
//	}
//
//	columnsStr := array.ToStringWithDelimiter(columns, ",")
//	query := strings.ReplaceAll(queryLogin, "${cols}", columnsStr)
//
//	rows, err := r.PgClient.Query(ctx, query, email)
//	if err != nil {
//		return nil, err
//	}
//
//	defer rows.Close()
//	for rows.Next() {
//		tempRes := &model.Users{}
//		err = rows.Scan(structs.GetAddressByFieldTag(tempRes, constants.DB_TAG, converts.AnyArrayToString(columns))...)
//		if err != nil {
//			return nil, err
//		}
//		result = append(result, tempRes)
//	}
//	err = rows.Err()
//	if err != nil {
//		return nil, err
//	}
//	return result, nil
//}

func (r *DatabaseRepository) GetUserByEmail(ctx context.Context, email string, columns ...any) (*model.Users, error) {
	var modelDb = new(model.Users)

	if len(columns) == 0 {
		columns = structs.GetColumns(&model.Users{})
	}

	columnsStr := array.ToStringWithDelimiter(columns, ",")
	query := strings.ReplaceAll(queryLogin, constants.DB_COLS, columnsStr)

	err := r.PgClient.QueryRow(ctx, query, email).Scan(structs.GetAddressByFieldTag(modelDb, constants.DB_TAG, converts.AnyArrayToString(columns))...)
	if err != nil {
		return nil, err
	}

	return modelDb, nil
}
