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
	CreatedDate     time.Time `db:"created_date"`
}

// constant for status
const (
	StatusDibayar      = "Sudah dibayar"
	StatusBelumDibayar = "Belum dibayar"
)

// constant for payment type
const (
	PaymentTypeCOD            = "COD"
	PaymentTypeManualTransfer = "Manual Transfer"
)

func (m Transactions) QueryInsert() string {
	return "INSERT INTO transactions (${cols}) VALUES ($1::bigint, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
}

func (m Transactions) QueryGet() string {
	return "SELECT ${cols} FROM transactions WHERE user_id=$1 ORDER BY id DESC"
}

func (m Transactions) QueryGetByID() string {
	return "SELECT ${cols} FROM transactions WHERE id=$1 ORDER BY id DESC"
}

func (m Transactions) QueryGetAll() string {
	return "SELECT ${cols} FROM transactions ORDER BY id DESC"
}
