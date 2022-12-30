package model

type Users struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	IsSeller bool   `db:"is_seller"`
}

func (m Users) QueryGetAll() string {
	return "SELECT ${cols} FROM users"
}

func (m Users) QueryGetByEmail() string {
	return "SELECT ${cols} FROM users WHERE email=$1"
}

func (m Users) QueryGetById() string {
	return "SELECT ${cols} FROM users WHERE id=$1"
}

func (m Users) QueryInsert() string {
	return "INSERT INTO users (${cols}) VALUES ($1, $2, $3, $4) RETURNING id"
}

func (m Users) QueryUpdate() string {
	return "UPDATE users SET ${cols} WHERE id = $1"
}
