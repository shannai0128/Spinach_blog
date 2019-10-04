package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/c479096292/Spinach_blog/config"
	"github.com/c479096292/Spinach_blog/model"
	"regexp"
)

type Article struct {
	ID              uint   `json:"aid"`
	Name			string
	Person_name         string     `json:"person_name"`
	Category_id 	   uint       `json:"category_id"`
	Title          string     `json:"title"`
	Summary		   string	   `json:"summary"`
	Content        string     `json:"content"`
	View_count     int        `json:"view_count"`
	Status         int  	  `json:"status"` // 0 正在审核, 1 过审  2 未过审
	Origin         int        `json:"origin"` //是否原创 1原创 0转载
	Violat_reason  sql.NullString     `json:"violat_reason"`
	Praise         int		  `json:"praise_count"` // 点赞
}

func (a *Article) Count() (int) {
	var article model.Article
	return article.GetTableTotal()
}

//  分页model
func (a *Article) ArticlesPaged(pageNum int, pageSize int) ([]*model.Article)  {
	var article model.Article
	articleObj, err := article.GetArticlesPage(pageNum, pageSize)
	if err != nil{
		config.Error(fmt.Sprintf("query article paged error: %s\n",err))
	}
	return articleObj
}

// 获取指定用户全部文章
func (a *Article) GetArticlesByPersonID(id int) ([]*model.Article)  {
	var article model.Article
	articleObj, err := article.GetArticlesByPersonID(id)
	if err != nil{
		config.Error(fmt.Sprintf("acquire user articles error: %s\n",err))
	}
	return articleObj
}

// 查找指定标题文章
func (a *Article) FindArticleByTitle(title string) (interface{})  {
	var article model.Article
	reg, _ :=regexp.Compile("^[a-zA-Z0-9\u4e00-\u9fa5]{2,16}$")
	ok := reg.MatchString(title)
	if !ok{
		return "please input a valid character"
	}
	articleObj, _ := article.FindArticleByTitle(title)
	return articleObj
}

// 新建文章
func (a *Article) InsertNewArticle() error {
	//var article model.Article
	for _, Sensitive := range config.ConfObj.SensitiveWords{
		if a.Title == Sensitive{
			return errors.New("含有敏感词汇,请重新检查")
		}
	}
	svcObj := model.Article{
		Category_id:a.Category_id,
		Content:a.Content,
		Title:a.Title,
		View_count:a.View_count,
		Person_name:a.Person_name,
		Summary:a.Summary,
		Origin:a.Origin,
	}
	svcObj.InsertNewArticle()

	return nil
}

// 修改文章
func (a *Article) EditArticle() error {
	//var article model.Article
	svcObj := model.Article{
		ID: a.ID,
		Category_id:a.Category_id,
		Content:a.Content,
		Title:a.Title,
		View_count:a.View_count,
		Person_name:a.Person_name,
		Summary:a.Summary,
		Origin:a.Origin,
	}
	svcObj.EditArticle()

	return nil
}

// 删除文章
func (a *Article) DelArticle(aid uint) error {
	svcObj := model.Article{
		ID:aid,
	}
	svcObj.DelArticle()
	return nil
}
