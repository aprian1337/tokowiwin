package model

type Snapshots struct {
	ID            int64  `db:"id"`
	TransactionID int64  `db:"transaction_id"`
	ProductPic    string `db:"product_pic"`
	ProductName   string `db:"product_name"`
	ProductPrice  int    `db:"product_price"`
	Qty           int64  `db:"qty"`
}

func (m Snapshots) QueryInsert() string {
	return "INSERT INTO snapshots (${cols}) VALUES ($1, $2, $3, $4, $5)"
}
func (m Snapshots) QueryGet() string {
	return "SELECT ${cols} FROM snapshots WHERE transaction_id=$1"
}

func (m Snapshots) QueryGetByIDs() string {
	return "SELECT ${cols} FROM snapshots WHERE id=ANY($1::int[])"
}
