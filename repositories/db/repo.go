package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/lib/pq"
	"log"
	"strings"
	"tokowiwin/constants"
	"tokowiwin/repositories/model"
	"tokowiwin/utils/array"
	"tokowiwin/utils/converts"
	"tokowiwin/utils/structs"
)

type RepositoryI interface {
	InsertUser(ctx context.Context, tx pgx.Tx, users *model.Users) (int64, error)
	GetUserByEmail(ctx context.Context, email string, columns ...any) (*model.Users, error)
	GetUserById(ctx context.Context, id int64, columns ...any) (*model.Users, error)
	GetUserAll(ctx context.Context, columns ...any) ([]*model.Users, error)
	UpdateUser(ctx context.Context, tx pgx.Tx, user *model.Users) error

	GetSnapshot(ctx context.Context, txId int64, columns ...any) ([]*model.Snapshots, error)
	InsertSnapshot(ctx context.Context, tx pgx.Tx, snapshot *model.Snapshots) error
	GetSnapshotsByIDs(ctx context.Context, ids []int64, columns ...any) ([]*model.Snapshots, error)
	GetSnapshotsByIDsMapped(ctx context.Context, ids []int64) (map[int64][]*model.Snapshots, error)

	InsertTransaction(ctx context.Context, tx pgx.Tx, transaction *model.Transactions) (int64, error)
	GetTransaction(ctx context.Context, userId int64, columns ...any) ([]*model.Transactions, error)
	GetTransactionByID(ctx context.Context, id int64, columns ...any) ([]*model.Transactions, error)
	GetAllTransactions(ctx context.Context, columns ...any) ([]*model.Transactions, error)

	GetCart(ctx context.Context, userId int64, columns ...any) ([]*model.Carts, error)
	InsertCart(ctx context.Context, tx pgx.Tx, cart *model.Carts) error
	UpdateCart(ctx context.Context, tx pgx.Tx, cart *model.Carts) error
	DeleteCart(ctx context.Context, tx pgx.Tx, id int64, userId int64) error
	DeleteAllCartByUserID(ctx context.Context, tx pgx.Tx, userId int64) error

	GetProductsAll(ctx context.Context, columns ...any) ([]*model.Products, error)
	GetProductsByID(ctx context.Context, id int64, columns ...any) (*model.Products, error)
	GetProductsByIDs(ctx context.Context, ids []int64, columns ...any) ([]*model.Products, error)
	GetProductsByIDsMapped(ctx context.Context, ids []int64) (map[int64]*model.Products, error)
	InsertProduct(ctx context.Context, tx pgx.Tx, products *model.Products) error
	DeleteProduct(ctx context.Context, tx pgx.Tx, id int64) error
	UpdateProduct(ctx context.Context, tx pgx.Tx, product *model.Products) error
}

func (r *DatabaseRepository) GetSnapshotsByIDs(ctx context.Context, ids []int64, columns ...any) ([]*model.Snapshots, error) {
	var (
		err     error
		result  = make([]*model.Snapshots, 0)
		modelDb = &model.Snapshots{}
	)

	if len(columns) == 0 {
		columns = structs.GetColumns(modelDb)
	}

	columnsStr := array.ToStringWithDelimiter(columns, ",")
	query := strings.ReplaceAll(modelDb.QueryGetByTransactionIDs(), constants.DbCols, columnsStr)

	rows, err := r.pgClient.Query(ctx, query, pq.Array(ids))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var temp = new(model.Snapshots)

		err = rows.Scan(structs.GetAddressByFieldTag(temp, constants.DbTag, converts.AnyArrayToString(columns))...)
		if err != nil {
			log.Default().Println(fmt.Sprintf("error=%v, query=%v", err, query))
			return nil, err
		}
		result = append(result, temp)
	}
	if err = rows.Err(); err != nil {
		log.Default().Println(fmt.Sprintf("error=%v, query=%v", err, query))
		return nil, err
	}

	return result, nil
}

func (r *DatabaseRepository) GetSnapshotsByIDsMapped(ctx context.Context, ids []int64) (map[int64][]*model.Snapshots, error) {
	d, err := r.GetSnapshotsByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	var result = make(map[int64][]*model.Snapshots, 0)
	for _, v := range d {
		result[v.TransactionID] = append(result[v.TransactionID], v)
	}

	return result, nil
}

func (r *DatabaseRepository) GetSnapshot(ctx context.Context, txId int64, columns ...any) ([]*model.Snapshots, error) {
	var (
		err     error
		result  = make([]*model.Snapshots, 0)
		modelDb = &model.Snapshots{}
	)

	if len(columns) == 0 {
		columns = structs.GetColumns(modelDb)
	}

	columnsStr := array.ToStringWithDelimiter(columns, ",")
	query := strings.ReplaceAll(modelDb.QueryGet(), constants.DbCols, columnsStr)

	rows, _ := r.pgClient.Query(ctx, query, txId)

	defer rows.Close()

	for rows.Next() {
		var temp = new(model.Snapshots)

		err = rows.Scan(structs.GetAddressByFieldTag(temp, constants.DbTag, converts.AnyArrayToString(columns))...)
		if err != nil {
			log.Default().Println(fmt.Sprintf("error=%v, query=%v", err, query))
			return nil, err
		}
		result = append(result, temp)
	}
	if err = rows.Err(); err != nil {
		log.Default().Println(fmt.Sprintf("error=%v, query=%v", err, query))
		return nil, err
	}

	return result, nil
}

func (r *DatabaseRepository) GetTransaction(ctx context.Context, userId int64, columns ...any) ([]*model.Transactions, error) {
	var (
		err     error
		result  = make([]*model.Transactions, 0)
		modelDb = &model.Transactions{}
	)

	if len(columns) == 0 {
		columns = structs.GetColumns(modelDb)
	}

	columnsStr := array.ToStringWithDelimiter(columns, ",")
	query := strings.ReplaceAll(modelDb.QueryGet(), constants.DbCols, columnsStr)

	rows, _ := r.pgClient.Query(ctx, query, userId)

	defer rows.Close()

	for rows.Next() {
		var temp = new(model.Transactions)

		err = rows.Scan(structs.GetAddressByFieldTag(temp, constants.DbTag, converts.AnyArrayToString(columns))...)
		if err != nil {
			log.Default().Println(fmt.Sprintf("error=%v, query=%v", err, query))
			return nil, err
		}
		result = append(result, temp)
	}
	if err = rows.Err(); err != nil {
		log.Default().Println(fmt.Sprintf("error=%v, query=%v", err, query))
		return nil, err
	}

	return result, nil
}

func (r *DatabaseRepository) GetTransactionByID(ctx context.Context, id int64, columns ...any) ([]*model.Transactions, error) {
	var (
		err     error
		result  = make([]*model.Transactions, 0)
		modelDb = &model.Transactions{}
	)

	if len(columns) == 0 {
		columns = structs.GetColumns(modelDb)
	}

	columnsStr := array.ToStringWithDelimiter(columns, ",")
	query := strings.ReplaceAll(modelDb.QueryGetByID(), constants.DbCols, columnsStr)

	rows, _ := r.pgClient.Query(ctx, query, id)

	defer rows.Close()

	for rows.Next() {
		var temp = new(model.Transactions)

		err = rows.Scan(structs.GetAddressByFieldTag(temp, constants.DbTag, converts.AnyArrayToString(columns))...)
		if err != nil {
			log.Default().Println(fmt.Sprintf("error=%v, query=%v", err, query))
			return nil, err
		}
		result = append(result, temp)
	}
	if err = rows.Err(); err != nil {
		log.Default().Println(fmt.Sprintf("error=%v, query=%v", err, query))
		return nil, err
	}

	return result, nil
}

func (r *DatabaseRepository) GetAllTransactions(ctx context.Context, columns ...any) ([]*model.Transactions, error) {
	var (
		err     error
		result  = make([]*model.Transactions, 0)
		modelDb = &model.Transactions{}
	)

	if len(columns) == 0 {
		columns = structs.GetColumns(modelDb)
	}

	columnsStr := array.ToStringWithDelimiter(columns, ",")
	query := strings.ReplaceAll(modelDb.QueryGetAll(), constants.DbCols, columnsStr)

	rows, _ := r.pgClient.Query(ctx, query)

	defer rows.Close()

	for rows.Next() {
		var temp = new(model.Transactions)

		err = rows.Scan(structs.GetAddressByFieldTag(temp, constants.DbTag, converts.AnyArrayToString(columns))...)
		if err != nil {
			log.Default().Println(fmt.Sprintf("error=%v, query=%v", err, query))
			return nil, err
		}
		result = append(result, temp)
	}
	if err = rows.Err(); err != nil {
		log.Default().Println(fmt.Sprintf("error=%v, query=%v", err, query))
		return nil, err
	}

	return result, nil
}

func (r *DatabaseRepository) GetCart(ctx context.Context, userId int64, columns ...any) ([]*model.Carts, error) {
	var (
		err     error
		result  = make([]*model.Carts, 0)
		modelDb = &model.Carts{}
	)

	if len(columns) == 0 {
		columns = structs.GetColumns(modelDb)
	}

	columnsStr := array.ToStringWithDelimiter(columns, ",")
	query := strings.ReplaceAll(modelDb.QueryGet(), constants.DbCols, columnsStr)

	rows, _ := r.pgClient.Query(ctx, query, userId)

	defer rows.Close()

	for rows.Next() {
		var temp = new(model.Carts)

		err = rows.Scan(structs.GetAddressByFieldTag(temp, constants.DbTag, converts.AnyArrayToString(columns))...)
		if err != nil {
			log.Default().Println(fmt.Sprintf("error=%v, query=%v", err, query))
			return nil, err
		}
		result = append(result, temp)
	}
	if err = rows.Err(); err != nil {
		log.Default().Println(fmt.Sprintf("error=%v, query=%v", err, query))
		return nil, err
	}

	return result, nil
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
	query := strings.ReplaceAll(modelDb.QueryGetAll(), constants.DbCols, columnsStr)

	rows, _ := r.pgClient.Query(ctx, query)

	defer rows.Close()

	for rows.Next() {
		var temp = new(model.Products)

		err = rows.Scan(structs.GetAddressByFieldTag(temp, constants.DbTag, converts.AnyArrayToString(columns))...)
		if err != nil {
			log.Default().Println(fmt.Sprintf("error=%v, query=%v", err, query))
			return nil, err
		}
		result = append(result, temp)
	}
	if err = rows.Err(); err != nil {
		log.Default().Println(fmt.Sprintf("error=%v, query=%v", err, query))
		return nil, err
	}

	return result, nil
}

func (r *DatabaseRepository) GetProductsByIDs(ctx context.Context, ids []int64, columns ...any) ([]*model.Products, error) {
	var (
		err     error
		result  = make([]*model.Products, 0)
		modelDb = &model.Products{}
	)

	if len(columns) == 0 {
		columns = structs.GetColumns(modelDb)
	}

	columnsStr := array.ToStringWithDelimiter(columns, ",")
	query := strings.ReplaceAll(modelDb.QueryGetByIDs(), constants.DbCols, columnsStr)

	rows, err := r.pgClient.Query(ctx, query, pq.Array(ids))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var temp = new(model.Products)

		err = rows.Scan(structs.GetAddressByFieldTag(temp, constants.DbTag, converts.AnyArrayToString(columns))...)
		if err != nil {
			log.Default().Println(fmt.Sprintf("error=%v, query=%v", err, query))
			return nil, err
		}
		result = append(result, temp)
	}
	if err = rows.Err(); err != nil {
		log.Default().Println(fmt.Sprintf("error=%v, query=%v", err, query))
		return nil, err
	}

	return result, nil
}

func (r *DatabaseRepository) GetProductsByIDsMapped(ctx context.Context, ids []int64) (map[int64]*model.Products, error) {
	d, err := r.GetProductsByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	var result = make(map[int64]*model.Products, 0)
	for _, v := range d {
		result[v.ID] = v
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
	query := strings.ReplaceAll(modelDb.QueryGetByID(), constants.DbCols, columnsStr)

	err = r.pgClient.QueryRow(ctx, query, id).Scan(structs.GetAddressByFieldTag(modelDb, constants.DbTag, converts.AnyArrayToString(columns))...)
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
	query := strings.ReplaceAll(modelDb.QueryInsert(), constants.DbCols, columnsStr)
	_, err = tx.Exec(ctx, query, structs.GetAddressByFieldTag(users, constants.DbTag, converts.AnyArrayToString(columns))...)
	if err != nil {
		return err
	}

	return nil
}

func (r *DatabaseRepository) DeleteCart(ctx context.Context, tx pgx.Tx, id int64, userId int64) error {
	var (
		err error
		m   = model.Carts{}
	)

	if tx == nil {
		tx, err = r.pgClient.Begin(ctx)
		if err != nil {
			return err
		}
	}

	_, err = tx.Exec(ctx, m.QueryDelete(), id, userId)
	if err != nil {
		return err
	}

	return nil
}

func (r *DatabaseRepository) DeleteAllCartByUserID(ctx context.Context, tx pgx.Tx, userId int64) error {
	var (
		err error
		m   = model.Carts{}
	)

	if tx == nil {
		tx, err = r.pgClient.Begin(ctx)
		if err != nil {
			return err
		}
	}

	_, err = tx.Exec(ctx, m.QueryDeleteByUserID(), userId)
	if err != nil {
		return err
	}

	return nil
}

func (r *DatabaseRepository) DeleteProduct(ctx context.Context, tx pgx.Tx, id int64) error {
	var (
		err error
		m   = model.Products{}
	)

	if tx == nil {
		tx, err = r.pgClient.Begin(ctx)
		if err != nil {
			return err
		}
	}

	_, err = tx.Exec(ctx, m.QueryDelete(), id)
	if err != nil {
		return err
	}

	return nil
}

func (r *DatabaseRepository) UpdateProduct(ctx context.Context, tx pgx.Tx, product *model.Products) error {
	var (
		err error
	)

	if tx == nil {
		tx, err = r.pgClient.Begin(ctx)
		if err != nil {
			return err
		}
	}

	queryVal := GetUpdateQuery(product, "id")
	query := strings.ReplaceAll(product.QueryUpdate(), constants.DbCols, queryVal)
	_, err = tx.Exec(ctx, query, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *DatabaseRepository) UpdateCart(ctx context.Context, tx pgx.Tx, cart *model.Carts) error {
	var (
		err error
	)

	if tx == nil {
		tx, err = r.pgClient.Begin(ctx)
		if err != nil {
			return err
		}
	}

	queryVal := GetUpdateQuery(cart, "user_id", "product_id")
	query := strings.ReplaceAll(cart.QueryUpdate(), constants.DbCols, queryVal)
	fmt.Println("QUERY", query)
	_, err = tx.Exec(ctx, query, cart.ProductID)
	if err != nil {
		return err
	}

	return nil
}

func (r *DatabaseRepository) GetUserAll(ctx context.Context, columns ...any) ([]*model.Users, error) {
	var (
		err     error
		result  = make([]*model.Users, 0)
		modelDb = &model.Users{}
	)

	if len(columns) == 0 {
		columns = structs.GetColumns(modelDb)
	}

	columnsStr := array.ToStringWithDelimiter(columns, ",")
	query := strings.ReplaceAll(modelDb.QueryGetAll(), constants.DbCols, columnsStr)

	rows, _ := r.pgClient.Query(ctx, query)

	defer rows.Close()

	for rows.Next() {
		var temp = new(model.Users)

		err = rows.Scan(structs.GetAddressByFieldTag(temp, constants.DbTag, converts.AnyArrayToString(columns))...)
		if err != nil {
			log.Default().Println(fmt.Sprintf("error=%v, query=%v", err, query))
			return nil, err
		}
		result = append(result, temp)
	}
	if err = rows.Err(); err != nil {
		log.Default().Println(fmt.Sprintf("error=%v, query=%v", err, query))
		return nil, err
	}

	return result, nil
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
	query := strings.ReplaceAll(modelDb.QueryGetByEmail(), constants.DbCols, columnsStr)

	err = r.pgClient.QueryRow(ctx, query, email).Scan(structs.GetAddressByFieldTag(modelDb, constants.DbTag, converts.AnyArrayToString(columns))...)
	if err != nil {
		return nil, err
	}

	return modelDb, nil
}

func (r *DatabaseRepository) GetUserById(ctx context.Context, id int64, columns ...any) (*model.Users, error) {
	var (
		err     error
		modelDb = new(model.Users)
	)

	if len(columns) == 0 {
		columns = structs.GetColumns(modelDb)
	}

	columnsStr := array.ToStringWithDelimiter(columns, ",")
	query := strings.ReplaceAll(modelDb.QueryGetById(), constants.DbCols, columnsStr)

	err = r.pgClient.QueryRow(ctx, query, id).Scan(structs.GetAddressByFieldTag(modelDb, constants.DbTag, converts.AnyArrayToString(columns))...)
	if err != nil {
		return nil, err
	}

	return modelDb, nil
}

func (r *DatabaseRepository) InsertUser(ctx context.Context, tx pgx.Tx, users *model.Users) (int64, error) {
	var (
		err     error
		id      int64
		modelDb = new(model.Users)
	)

	if tx == nil {
		tx, err = r.pgClient.Begin(ctx)
		if err != nil {
			return 0, err
		}
	}
	columns := structs.GetColumns(modelDb, "id")
	columnsStr := array.ToStringWithDelimiter(columns, ",")
	query := strings.ReplaceAll(modelDb.QueryInsert(), constants.DbCols, columnsStr)
	err = tx.QueryRow(ctx, query, structs.GetAddressByFieldTag(users, constants.DbTag, converts.AnyArrayToString(columns))...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *DatabaseRepository) UpdateUser(ctx context.Context, tx pgx.Tx, user *model.Users) error {
	var (
		err error
	)

	if tx == nil {
		tx, err = r.pgClient.Begin(ctx)
		if err != nil {
			return err
		}
	}

	queryVal := GetUpdateQuery(user, "id", "name", "email", "is_seller")
	query := strings.ReplaceAll(user.QueryUpdate(), constants.DbCols, queryVal)
	_, err = tx.Exec(ctx, query, user.ID)
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
	query := strings.ReplaceAll(modelDb.QueryInsert(), constants.DbCols, columnsStr)
	_, err = tx.Exec(ctx, query, structs.GetAddressByFieldTag(snapshot, constants.DbTag, converts.AnyArrayToString(columns))...)
	if err != nil {
		return err
	}

	return nil
}

func (r *DatabaseRepository) InsertTransaction(ctx context.Context, tx pgx.Tx, transaction *model.Transactions) (int64, error) {
	var (
		err     error
		modelDb = new(model.Transactions)
		id      int64
	)

	if tx == nil {
		tx, err = r.pgClient.Begin(ctx)
		if err != nil {
			return id, err
		}
	}

	columns := structs.GetColumns(modelDb)
	columnsStr := array.ToStringWithDelimiter(columns, ",")
	query := strings.ReplaceAll(modelDb.QueryInsert(), constants.DbCols, columnsStr)
	fmt.Println("QUERY", query)
	err = tx.QueryRow(ctx, query, structs.GetAddressByFieldTag(transaction, constants.DbTag, converts.AnyArrayToString(columns))...).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
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
	query := strings.ReplaceAll(modelDb.QueryInsert(), constants.DbCols, columnsStr)
	_, err = tx.Exec(ctx, query, structs.GetAddressByFieldTag(cart, constants.DbTag, converts.AnyArrayToString(columns))...)
	if err != nil {
		return err
	}

	return nil
}
