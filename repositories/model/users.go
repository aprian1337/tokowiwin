package model

type Users struct {
	ID       int64  `db:"id" autoinc:"true"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
