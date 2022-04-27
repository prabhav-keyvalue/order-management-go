package entity

type Customer struct {
	BaseEntity
	Name       string `json:"name" example:"max"`
	Phone      string `json:"phone" example:"+9876467658"`
	Street     string `json:"street" example:"olive"`
	State      string `json:"state" example:"kerala"`
	City       string `json:"city" example:"gotham"`
	ProfilePic string `json:"profilePic" example:"https://picsum.photos/200"`
}
