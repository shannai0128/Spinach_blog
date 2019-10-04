package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

//type BaseModel struct {
//	Create_time time.Time `json:"create_time"`
//	Update_time time.Time `json:"updated_time"`
//}
type BaseModel struct {
	ID        uint      `gorm:"primary_key"` // 自增
	CreatedAt time.Time `gorm:"not null" json:"-"`
	UpdatedAt time.Time `gorm:"not null" json:"-"`
}

type Where func(*gorm.DB) *gorm.DB
