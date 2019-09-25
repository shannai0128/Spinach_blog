package model

import (
	"database/sql"
	"fmt"
	"github.com/c479096292/Spinach_blog/config"
	"github.com/c479096292/Spinach_blog/db"
)

type Article struct {
	ID             uint       `json:"aid"`
	//CreatedAt      time.Time  `json:"-"`
	//UpdatedAt      time.Time  `json:"-"`
	BaseModel
	//DeletedAt      *time.Time `sql:"index" json:"-"`
	Person_name         uint     `json:"Person_name"`
	CategoryID 	   uint       `json:"cid"`
	Title          string     `json:"title"`
	Summary		   string	   `json:"summary"`
	Content        string     `json:"content"`
	View_count     int        `json:"view_count"`
	Status         int  	  `json:"status"` // 0 未过审, 1 过审
	Origin         int        `json:"origin"` //是否原创 1原创 0转载
	Violat_reason  sql.NullString     `json:"violat_reason"`
	//Comment_id       string     `json:"comment_id"`
	//Comment_count  int    	  `json:"comment_count"`
	Praise_count         int		  `json:"praise_count"` // 点赞
}

// 获取表数据量
func GetTableTotal(tableName string) (total int) {
	sqlstr := "select count(*) from " + tableName
	err := db.DB.Get(&total, sqlstr)
	if err != nil{
		err_info :=fmt.Sprintf("Get table %s Total failed, err:%v\n", tableName, err)
		 config.Error(err_info)
	}
	return
}

// 分页
func GetArticlesPage(pageNum int, pageSize int) ([]*Article, error) {
	var articles []*Article
	sqlstr := "select * from article limit (?-1)*?,?;"
	err := db.DB.Select(&articles,sqlstr,pageSize,pageSize)
	if err != nil{
		err_info :=fmt.Sprintf("Get article %s paged failed, err:%v\n", err)
		config.Error(err_info)
	}

	return articles, nil
}

// 获取指定用户全部文章
func GetArticlesByPersonID(id int) ([]*Article, error) {
	var articles []*Article
	sqlstr := "select * from article where person_name = (select person_name from person where id=?);"
	err := db.DB.Select(&articles,sqlstr,id)
	if err != nil{
		err_info :=fmt.Sprintf("Get article %s paged failed, err:%v\n", err)
		config.Error(err_info)
	}
	return articles, nil
}



