package entity

type Customer struct {
	BaseEntity
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Street     string `json:"street"`
	State      string `json:"state"`
	City       string `json:"city"`
	ProfilePic string `json:"profilePic"`
}
