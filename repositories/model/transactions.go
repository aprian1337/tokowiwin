package model

type Transactions struct {
	ID              int64  `db:"id"`
	UserID          int64  `db:"user_id"`
	ReceiverName    string `db:"receiver_name"`
	ReceiverPhone   string `db:"receiver_phone"`
	ReceiverAddress string `db:"receiver_address"`
	PaymentType     string `db:"payment_type"`
}

func (m Transactions) QueryInsert() string {
	return "INSERT INTO transactions (${cols}) VALUES ($1, $2, $3, $4, $5)"
}
