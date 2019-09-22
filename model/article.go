package model

import (
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
	Person_id         uint     `json:"Person_id"`
	CategoryID 	   uint       `json:"cid"`
	Title          string     `json:"title"`
	Summary		   string	   `json:"summary"`
	Content        string     `json:"content"`
	View_count     int        `json:"view_count"`
	//img          int        `json:"img"`
	Origin         int        `json:"origin"` //是否原创 1原创 0转载
	State          int        `json:"state"`     //0正常发布 2并未发布(草稿箱)
	Violat_reason  string     `json:"violat_reason"`
	praise         int		  `json:"praise"`
	Comment_id       string     `json:"comment_id"`
	Comment_count  int    	  `json:"comment_count"`
	Praise         int		  `json:"praise"`
}

// 获取表数据量
func GetTableTotal(tableName string) (total int) {
	sqlstr := "select count(*) from ?"
	err := db.DB.QueryRow(sqlstr,tableName).Scan(&total)
	if err != nil{
		err_info :=fmt.Sprintf("Get table %s Total failed, err:%v\n", tableName, err)
		 config.Error(err_info)
	}
	return
}

func GetArticlesPage(pageNum int, pageSize int) ([]*Article, error) {
	var articles []*Article
	sqlstr := "select * from article limit (?-1)*?,?"
	err := db.DB.Select(articles,sqlstr,pageSize,pageSize)
	if err != nil{
		err_info :=fmt.Sprintf("Get article %s paged failed, err:%v\n", err)
		config.Error(err_info)
	}

	return articles, nil
}


