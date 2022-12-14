package model

import "database/sql"

type Carts struct {
	ProductID sql.NullInt64 `db:"product_id"`
	UserID    int64         `db:"user_id"`
	Qty       sql.NullInt64 `db:"qty"`
}

func (m Carts) QueryGetByUserID() string {
	return "SELECT ${cols} FROM carts WHERE user_id=$1"
}

func (m Carts) QueryInsert() string {
	return "INSERT INTO carts (${cols}) VALUES ($1, $2, $3)"
}

func (m Carts) QueryUpdate() string {
	return "UPDATE carts SET ${cols} WHERE product_id = $1"
}

func (m Carts) QueryGet() string {
	return "SELECT ${cols} FROM carts WHERE user_id = $1"
}

func (m Carts) QueryDelete() string {
	return "UPDATE products SET ${cols} WHERE product_id = $1"
}
