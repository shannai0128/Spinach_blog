package model

type Category struct {
	ID uint  `gorm:"primary_key"`
	Name string  `gorm:"type:varchar(20);not null"`
}
