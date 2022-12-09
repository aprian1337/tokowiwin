package queries

const (
	QueryUsersByEmail = "SELECT ${cols} FROM users WHERE email=$1"
	QueryInsertUsers  = "INSERT INTO users (${cols}) VALUES ($1, $2, $3)"
)
