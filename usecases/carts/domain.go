package carts

type Carts struct {
	ID           int64  `json:"id"`
	ProductName  string `json:"product_name"`
	ProductStock int    `json:"product_stock"`
	ProductPrice string `json:"product_price"`
	ProductPic   string `json:"product_pic"`
	Qty          int64  `json:"qty"`
}
