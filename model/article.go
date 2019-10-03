package model

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/c479096292/Spinach_blog/config"
	"github.com/c479096292/Spinach_blog/db"
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	ID    uint  `gorm:"primary_key"`
	Person_name         string     `gorm:"type:varchar(256);not null"`
	Category_id 	   uint       `gorm:"type:bigint(20);not null;unique_index" json:"category_id"`
	Title          string     `gorm:"type:varchar(255);not null"`
	Summary		   string	   `gorm:"type:varchar(255);not null"`
	Content        string     `gorm:"type:longtext;not null"`
	View_count     int        `gorm:"type:int;index:idx_view_count"`
	Status         int  	  `gorm:"type:int;default:0"` // 0 正在审核, 1 过审  2 未过审
	Origin         int        `gorm:"type:int"` //是否原创 1原创 0转载
	Violat_reason  sql.NullString     `gorm:"type:text"`
	//Comment_id       string     `json:"comment_id"`
	//Comment_count  int    	  `json:"comment_count"`
	CreateTime     time.Time  `gorm:"not null" json:"create_time"`
	UpdateTime     time.Time  `gorm:"not null" json:"update_time"`
	Praise         int		  `gorm:"type:bigint(10);default:0"` // 点赞
}

func (a *Article) BeforeUpdate() (err error) {
	if (len(a.Title) > 18 ) {
		err = errors.New("article title is already greater than 18 character")
		config.Warn(err)
	}
	return
}

func (a *Article) db() *gorm.DB {
	return db.GetOrm()
}

// 获取表数据量
func (a *Article) GetTableTotal() int {
	var v interface{}
	err := a.db().Find(&v,&Article{}).Error
	if err !=nil {
		err_info :=fmt.Sprintf("Get article Total failed, err:%v\n", err)
		config.Error(err_info)
	}
	val, ok := v.(int);
	if  !ok {
		err_info :="result is not type of int"
		config.Error(err_info)
		panic(err_info)
	}
	return val
}

func GetTableTotal(tableName string) (total int) {
	sqlstr := "select count(*) from " + tableName
	err := db.DB.Get(&total, sqlstr)
	if err != nil{
		err_info :=fmt.Sprintf("Get table %s Total failed, err:%v\n", tableName, err)
		 config.Error(err_info)
		panic(err)
	}
	return
}

// 分页
func GetArticlesPage(pageNum int, pageSize int) ([]*Article, error) {
	var articles []*Article
	args1 := (pageNum-1) * pageSize
	sqlstr := "select * from article limit ?,?;"
	err := db.DB.Select(&articles,sqlstr,args1, pageSize)
	if err != nil{
		err_info :=fmt.Sprintf("Get article page failed, err:%s\n", err)
		config.Error(err_info)
		panic(err)
	}
	return articles, nil
}

// 获取指定用户全部文章
func GetArticlesByPersonID(id int) ([]*Article, error) {
	var articles []*Article
	sqlstr := "select * from article where person_name = (select person_name from person where id=?);"
	err := db.DB.Select(&articles,sqlstr,id)
	if err != nil{
		err_info :=fmt.Sprintf("Get user articles failed, err:%s\n", err)
		config.Error(err_info)
		panic(err)
	}
	return articles, nil
}

// 查找文章by title
func FindArticleByTitle(title string) ([]*Article, error) {
	var articles []*Article
	sqlstr := "select * from article where title like '%"+title+"%';"
	err := db.DB.Select(&articles,sqlstr)
	if err != nil{
		err_info :=fmt.Sprintf("find the %s article failed, err:%s\n", title, err)
		config.Error(err_info)
		panic(err)
	}
	return articles, nil
}

// 添加新文章
//func InsertNewArticle(a Article) (error) {
//	sqlstr := "insert into article(category_id,content,title,view_count,person_name,summary,origin) value(?,?,?,?,?,?,?);"
//	_ , err := db.DB.Exec(sqlstr, a.Category_id, a.Content, a.Title,a.View_count, a.Person_name, a.Summary, a.Origin)
//	if err != nil{
//		err_info :=fmt.Sprintf("create new article failed, err:%s\n", err)
//		config.Error(err_info)
//		panic(err)
//	}
//	return nil
//}
//
//// 修改文章
//func EditArticle(a Article) error {
//	sqlstr := "update article set category_id=?,content=?,title=?,view_count=?,person_name=?,summary=?,origin=? where id=?;"
//	_ , err := db.DB.Exec(sqlstr, a.Category_id, a.Content, a.Title,a.View_count, a.Person_name, a.Summary, a.Origin, a.ID)
//	if err != nil{
//		err_info :=fmt.Sprintf("edit the article failed, err:%s\n", err)
//		config.Error(err_info)
//		panic(err)
//	}
//	return nil
//}
//
//func DelArticle(aid int) error {
//	sqlstr := "delete from article where id=?;"
//	_ , err := db.DB.Exec(sqlstr, aid)
//	if err != nil{
//		err_info :=fmt.Sprintf("delete article failed, err:%s\n", err)
//		config.Error(err_info)
//		panic(err)
//	}
//	return nil
//}