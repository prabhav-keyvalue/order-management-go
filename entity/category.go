package entity

type Category struct {
	BaseEntity
	Name           string `json:"name" example:"TV"`
	ParentCategory string `json:"parentCategory" example:"Electronics"`
}
