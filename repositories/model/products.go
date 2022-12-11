package model

type Products struct {
	ID           int64  `db:"id"`
	ProductName  string `db:"product_name"`
	ProductStock int    `db:"product_stock"`
	ProductPrice int    `db:"product_price"`
	ProductPic   string `db:"product_pic"`
}

func (m Products) QueryGetAll() string {
	return "SELECT ${cols} FROM products"
}

func (m Products) QueryGetByID() string {
	return "SELECT ${cols} FROM products WHERE id=$1"
}

func (m Products) QueryInsert() string {
	return "INSERT INTO products (${cols}) VALUES ($1, $2, $3, $4)"
}
