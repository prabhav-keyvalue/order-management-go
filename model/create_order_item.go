package model

type CreateOrderItemInput struct {
	OrderId   string  `json:"orderId"`
	ProductId string  `json:"productId"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	RowTotal  int     `json:"rowTotal"`
}
