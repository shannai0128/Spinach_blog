package model

import (
	"database/sql"
	"time"
)

type Person struct {
	BaseModel
	ID           int
	PersonName     string
	IdCard		string
	PersonInfoID   uint     `json:"person_infoid"`
	PassWord     string `json:"password"`        //sha1
	gender		 int
	article_id   int
	LoginIP      sql.NullString
	LoginTime    time.Time
	IsActive     int      `json:"is_active"` //0可用 1禁用
}
