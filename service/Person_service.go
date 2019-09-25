package service

import (
	"database/sql"
	"github.com/c479096292/Spinach_blog/model"
	"time"
)

type Person struct {
	ID            int
	Name 		  string // 表名
	PersonName     string
	IdCard		string
	PersonInfoID   uint
	PassWord     string
	gender		 int
	article_id   int
	LoginIP      sql.NullString
	LoginTime    time.Time
	IsActive     int
}

func (p *Person) Count() (int) {
	return model.GetTableTotal(p.Name)
}

