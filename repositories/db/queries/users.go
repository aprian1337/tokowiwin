package queries

const (
	//Query Select
	QueryUsersByEmail = "SELECT ${cols} FROM users WHERE email=$1"
	QueryProductsAll  = "SELECT ${cols} FROM products"
	QueryProductsByID = "SELECT ${cols} FROM products WHERE id=$1"

	//Query Insert
	QueryInsertUsers       = "INSERT INTO users (${cols}) VALUES ($1, $2, $3, $4)"
	QueryInsertProducts    = "INSERT INTO products (${cols}) VALUES ($1, $2, $3, $4)"
	QueryInsertSnapshot    = "INSERT INTO snapshots (${cols}) VALUES ($1, $2, $3, $4, $5)"
	QueryInsertTransaction = "INSERT INTO transactions (${cols}) VALUES ($1, $2, $3, $4, $5)"
	QueryInsertCart        = "INSERT INTO carts (${cols}) VALUES ($1, $2, $3)"
)
