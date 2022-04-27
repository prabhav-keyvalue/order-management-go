package entity

type Product struct {
	BaseEntity
	Name        string   `json:"name" example:"pepsi"`
	Image       string   `json:"image" example:"https://picsum.photos/200"`
	Description string   `json:"description" example:"test description"`
	UnitPrice   float64  `json:"unitPrice" example:"200"`
	CategoryId  string   `json:"categoryId" example:"951c54e9-4b64-42fe-9d56-e8a9babc3f89"`
	Category    Category `json:"category"`
}
