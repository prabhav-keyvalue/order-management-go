package entity

type Order struct {
	BaseEntity
	CustomerId    string      `json:"customerId" example:"0189fabc-1afc-49f9-bf68-95453466b50d"`
	Customer      *Customer   `json:"customer,omitempty"`
	TotalQuantity int         `json:"totalQuantity" example:"100"`
	TotalPrice    float64     `json:"totalPrice" example:"1000"`
	OrderItems    []OrderItem `json:"orderItems,omitempty"`
}
