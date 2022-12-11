package model

type Carts struct {
	ProductID int64 `db:"product_id"`
	UserID    int64 `db:"user_id"`
	Qty       int   `db:"qty"`
}

func (m Carts) QueryGetByUserID() string {
	return "SELECT ${cols} FROM carts WHERE user_id=$1"
}

func (m Carts) QueryInsert() string {
	return "INSERT INTO carts (${cols}) VALUES ($1, $2, $3)"
}
