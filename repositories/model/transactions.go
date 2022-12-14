package model

import "time"

type Transactions struct {
	ID              int64     `db:"id"`
	UserID          int64     `db:"user_id"`
	ReceiverName    string    `db:"receiver_name"`
	ReceiverPhone   string    `db:"receiver_phone"`
	ReceiverAddress string    `db:"receiver_address"`
	PaymentType     string    `db:"payment_type"`
	Status          string    `db:"status"`
	Date            time.Time `db:"date"`
}

const (
	StatusDibayar      = "Sudah dibayar"
	StatusBelumDibayar = "Belum dibayar"
)

func (m Transactions) QueryInsert() string {
	return "INSERT INTO transactions (${cols}) VALUES ($1, $2, $3, $4, $5) RETURNING id"
}

func (m Transactions) QueryGet() string {
	return "SELECT ${cols} FROM transactions WHERE user_id=$1"
}
