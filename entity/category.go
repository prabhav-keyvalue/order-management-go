package entity

type Category struct {
	BaseEntity
	Name           string `json:"name"`
	ParentCategory string `json:"parentCategory"`
}
