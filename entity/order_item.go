package entity

type OrderItem struct {
	BaseEntity
	OrderId   string  `json:"orderId"`
	Order     Order   `json:"order"`
	ProductId string  `json:"productId"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	RowTotal  float64 `json:"rowTotal"`
}
