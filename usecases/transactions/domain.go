package transactions

type TransactionList struct {
	TransactionID      int64              `json:"transaction_id"`
	ProductName        string             `json:"product_name"`
	AnotherProduct     string             `json:"another_product"`
	Date               string             `json:"date"`
	ProductPic         string             `json:"product_pic"`
	TransactionDetails *TransactionDetail `json:"transaction_details"`
}

type TransactionDetail struct {
	ReceiverName        string                `json:"receiver_name"`
	ReceiverPhone       string                `json:"receiver_phone"`
	ReceiverAddress     string                `json:"address"`
	PaymentType         string                `json:"payment_type"`
	TotalBill           string                `json:"total_bill"`
	TransactionProducts []*TransactionProduct `json:"transaction_products"`
}

type TransactionProduct struct {
	ProductName       string `json:"product_name"`
	ProductPic        string `json:"product_pic"`
	ProductQty        int64  `json:"product_qty"`
	ProductPrice      string `json:"product_price"`
	ProductTotalPrice string `json:"product_total_price"`
}
