package service

import (
	"github.com/c479096292/Spinach_blog/model"
)

type Article struct {
	ID            int
	Name 		  string // 表名
	TagID         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreatedBy     string
	ModifiedBy    string
	PageNum  int
	PageSize int
}

func (a *Article) Count() (int) {
	return model.GetTableTotal(a.Name)
}

//func (a *Article) GetArticlesyPersonID(id int) ([]*model.Article, err) {
//
//
//}
