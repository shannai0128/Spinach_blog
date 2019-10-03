package db

import (
	"fmt"
	"github.com/c479096292/Spinach_blog/model"
)

func InitRelations() {
	db := GetOrm()
	// 判断存不存在表, 不存在就新建, 否则就是自动迁移(其他修改)
	if !db.HasTable("blog_Person") {
		fmt.Println(1111111111111)
		db.CreateTable(&model.Person{}, &model.PersonInfo{}, &model.Category{}, &model.Article{})
	} else {
		fmt.Println(2222222222222)
		db.AutoMigrate(&model.Person{}, &model.PersonInfo{}, &model.Category{}, &model.Article{})
	}
}

