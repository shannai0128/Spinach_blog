package model

import (
	"github.com/c479096292/Spinach_blog/db"
)

func InitRelations() {
	dbobj := db.GetOrm()

	if !dbobj.HasTable("blog_person") {
		dbobj.CreateTable(&Person{}, &PersonInfo{}, &Category{}, &Article{})
	} else {

		dbobj.AutoMigrate(&Person{}, &PersonInfo{}, &Category{}, &Article{})
	}
}

