package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"strings"
	"tokowiwin/constants"
	"tokowiwin/repositories/model"
	"tokowiwin/utils/array"
	"tokowiwin/utils/converts"
	"tokowiwin/utils/structs"
)

type RepositoryI interface {
	GetUserByEmail(ctx context.Context, email string, columns ...any) (*model.Users, error)
	GetProductsAll(ctx context.Context, columns ...any) ([]*model.Products, error)
	GetProductsByID(ctx context.Context, id int64, columns ...any) (*model.Products, error)

	InsertUser(ctx context.Context, tx pgx.Tx, users *model.Users) error
	InsertProduct(ctx context.Context, tx pgx.Tx, products *model.Products) error
	InsertSnapshot(ctx context.Context, tx pgx.Tx, snapshot *model.Snapshots) error
	InsertTransaction(ctx context.Context, tx pgx.Tx, transaction *model.Transactions) error
	InsertCart(ctx context.Context, tx pgx.Tx, cart *model.Carts) error
}

func (r *DatabaseRepository) GetProductsAll(ctx context.Context, columns ...any) ([]*model.Products, error) {
	var (
		err     error
		result  = make([]*model.Products, 0)
		modelDb = &model.Products{}
	)

	if len(columns) == 0 {
		columns = structs.GetColumns(modelDb)
	}

	columnsStr := array.ToStringWithDelimiter(columns, ",")
	query := strings.ReplaceAll(modelDb.QueryGetAll(), constants.DB_COLS, columnsStr)

	rows, _ := r.pgClient.Query(ctx, query)

	defer rows.Close()

	for rows.Next() {
		var temp = new(model.Products)

		err = rows.Scan(structs.GetAddressByFieldTag(temp, constants.DB_TAG, converts.AnyArrayToString(columns))...)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, temp)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result, nil
}

func (r *DatabaseRepository) GetProductsByID(ctx context.Context, id int64, columns ...any) (*model.Products, error) {
	var (
		err     error
		modelDb = new(model.Products)
	)

	if len(columns) == 0 {
		columns = structs.GetColumns(modelDb)
	}

	columnsStr := array.ToStringWithDelimiter(columns, ",")
	query := strings.ReplaceAll(modelDb.QueryGetByID(), constants.DB_COLS, columnsStr)

	err = r.pgClient.QueryRow(ctx, query, id).Scan(structs.GetAddressByFieldTag(modelDb, constants.DB_TAG, converts.AnyArrayToString(columns))...)
	if err != nil {
		return nil, err
	}

	return modelDb, nil
}

func (r *DatabaseRepository) InsertProduct(ctx context.Context, tx pgx.Tx, users *model.Products) error {
	var (
		err     error
		modelDb = new(model.Products)
	)

	if tx == nil {
		tx, err = r.pgClient.Begin(ctx)
		if err != nil {
			return err
		}
	}

	columns := structs.GetColumns(modelDb, "id")
	columnsStr := array.ToStringWithDelimiter(columns, ",")
	query := strings.ReplaceAll(modelDb.QueryInsert(), constants.DB_COLS, columnsStr)
	_, err = tx.Exec(ctx, query, structs.GetAddressByFieldTag(users, constants.DB_TAG, converts.AnyArrayToString(columns))...)
	if err != nil {
		return err
	}

	return nil
}

func (r *DatabaseRepository) GetUserByEmail(ctx context.Context, email string, columns ...any) (*model.Users, error) {
	var (
		err     error
		modelDb = new(model.Users)
	)

	if len(columns) == 0 {
		columns = structs.GetColumns(modelDb)
	}

	columnsStr := array.ToStringWithDelimiter(columns, ",")
	query := strings.ReplaceAll(modelDb.QueryGetByEmail(), constants.DB_COLS, columnsStr)

	err = r.pgClient.QueryRow(ctx, query, email).Scan(structs.GetAddressByFieldTag(modelDb, constants.DB_TAG, converts.AnyArrayToString(columns))...)
	if err != nil {
		return nil, err
	}

	return modelDb, nil
}

func (r *DatabaseRepository) InsertUser(ctx context.Context, tx pgx.Tx, users *model.Users) error {
	var (
		err     error
		modelDb = new(model.Users)
	)

	if tx == nil {
		tx, err = r.pgClient.Begin(ctx)
		if err != nil {
			return err
		}
	}
	columns := structs.GetColumns(modelDb, "id")
	columnsStr := array.ToStringWithDelimiter(columns, ",")
	query := strings.ReplaceAll(modelDb.QueryInsert(), constants.DB_COLS, columnsStr)
	_, err = tx.Exec(ctx, query, structs.GetAddressByFieldTag(users, constants.DB_TAG, converts.AnyArrayToString(columns))...)
	if err != nil {
		return err
	}

	return nil
}

func (r *DatabaseRepository) InsertSnapshot(ctx context.Context, tx pgx.Tx, snapshot *model.Snapshots) error {
	var (
		err     error
		modelDb = new(model.Snapshots)
	)

	if tx == nil {
		tx, err = r.pgClient.Begin(ctx)
		if err != nil {
			return err
		}
	}

	columns := structs.GetColumns(modelDb, "id")
	columnsStr := array.ToStringWithDelimiter(columns, ",")
	query := strings.ReplaceAll(modelDb.QueryInsert(), constants.DB_COLS, columnsStr)
	_, err = tx.Exec(ctx, query, structs.GetAddressByFieldTag(snapshot, constants.DB_TAG, converts.AnyArrayToString(columns))...)
	if err != nil {
		return err
	}

	return nil
}

func (r *DatabaseRepository) InsertTransaction(ctx context.Context, tx pgx.Tx, transaction *model.Transactions) error {
	var (
		err     error
		modelDb = new(model.Transactions)
	)

	if tx == nil {
		tx, err = r.pgClient.Begin(ctx)
		if err != nil {
			return err
		}
	}

	columns := structs.GetColumns(modelDb, "id")
	columnsStr := array.ToStringWithDelimiter(columns, ",")
	query := strings.ReplaceAll(modelDb.QueryInsert(), constants.DB_COLS, columnsStr)
	_, err = tx.Exec(ctx, query, structs.GetAddressByFieldTag(transaction, constants.DB_TAG, converts.AnyArrayToString(columns))...)
	if err != nil {
		return err
	}

	return nil
}

func (r *DatabaseRepository) InsertCart(ctx context.Context, tx pgx.Tx, cart *model.Carts) error {
	var (
		err     error
		modelDb = new(model.Carts)
	)

	if tx == nil {
		tx, err = r.pgClient.Begin(ctx)
		if err != nil {
			return err
		}
	}

	columns := structs.GetColumns(modelDb, "id")
	columnsStr := array.ToStringWithDelimiter(columns, ",")
	query := strings.ReplaceAll(modelDb.QueryInsert(), constants.DB_COLS, columnsStr)
	_, err = tx.Exec(ctx, query, structs.GetAddressByFieldTag(cart, constants.DB_TAG, converts.AnyArrayToString(columns))...)
	if err != nil {
		return err
	}

	return nil
}
