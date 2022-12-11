package products

type Products struct {
	ID           int64  `json:"id"`
	ProductName  string `json:"product_name"`
	ProductStock string `json:"product_stock"`
	ProductPrice string `json:"product_price"`
	ProductPic   string `json:"product_pic"`
}
