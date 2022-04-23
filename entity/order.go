package entity

type Order struct {
	BaseEntity
	CustomerId    string      `json:"customerId"`
	Customer      Customer    `json:"customer"`
	TotalQuantity int         `json:"totalQuantity"`
	TotalPrice    float64     `json:"totalPrice"`
	OrderItems    []OrderItem `json:"orderItems"`
}
