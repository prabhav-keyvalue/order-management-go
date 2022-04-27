package entity

type OrderItem struct {
	BaseEntity
	OrderId   string  `json:"orderId" example:"0189fabc-1afc-49f9-bf68-95453466b50d"`
	ProductId string  `json:"productId" example:"951c54e9-4b64-42fe-9d56-e8a9babc3f89"`
	Quantity  int     `json:"quantity" example:"100"`
	Price     float64 `json:"price" example:"9.78"`
	RowTotal  float64 `json:"rowTotal" example:"978.00"`
}
