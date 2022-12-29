package model

type Products struct {
	ID           int64  `db:"id"`
	ProductName  string `db:"product_name"`
	ProductStock int    `db:"product_stock"`
	ProductPrice int    `db:"product_price"`
	ProductPic   string `db:"product_pic"`
}

func (m Products) QueryGetAll() string {
	return "SELECT ${cols} FROM products WHERE product_stock > 0"
}

func (m Products) QueryGetByID() string {
	return "SELECT ${cols} FROM products WHERE id=$1"
}

func (m Products) QueryGetByIDs() string {
	return "SELECT ${cols} FROM products WHERE id=ANY($1::int[])"
}

func (m Products) QueryInsert() string {
	return "INSERT INTO products (${cols}) VALUES ($1, $2, $3, $4)"
}

func (m Products) QueryUpdate() string {
	return "UPDATE products SET ${cols} WHERE id = $1"
}

func (m Products) QueryDelete() string {
	return "DELETE FROM products WHERE id = $1"
}
