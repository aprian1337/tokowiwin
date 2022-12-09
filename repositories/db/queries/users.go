package queries

const (
	QueryUsersByEmail = "SELECT ${cols} FROM users WHERE email=$1"
)
