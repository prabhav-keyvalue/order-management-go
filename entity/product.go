package entity

type Product struct {
	BaseEntity
	Name        string   `json:"name"`
	Image       string   `json:"image"`
	Description string   `json:"description"`
	UnitPrice   float64  `json:"unitPrice"`
	CategoryId  string   `json:"categoryId"`
	Category    Category `json:"category"`
}
