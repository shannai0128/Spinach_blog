package model

type Category struct {
	ID int
	BaseModel
	Name string `json:"name"`
}
