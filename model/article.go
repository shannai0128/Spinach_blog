package model

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/c479096292/Spinach_blog/config"
	"github.com/c479096292/Spinach_blog/db"
	"github.com/jinzhu/gorm"
)

type Article struct {
	ID    uint  `gorm:"primary_key"`
	Person_name         string     `gorm:"type:varchar(256);not null"`
	Category_id 	   uint       `gorm:"type:bigint(20);not null" json:"category_id"`
	Title          string     `gorm:"type:varchar(255);not null"`
	Summary		   string	   `gorm:"type:varchar(255);not null"`
	Content        string     `gorm:"type:longtext;not null"`
	View_count     int        `gorm:"type:int;index:idx_view_count;default:0"`
	Status         int  	  `gorm:"type:int;default:0"` // 0 正在审核, 1 过审  2 未过审
	Origin         int        `gorm:"type:int"` //是否原创 1原创 0转载
	Violat_reason  sql.NullString     `gorm:"type:text"`
	//Comment_id       string     `json:"comment_id"`
	//Comment_count  int    	  `json:"comment_count"`
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
	var v int
	err := a.db().Model(&Article{}).Count(&v).Error
	if err !=nil {
		err_info :=fmt.Sprintf("Get article Total failed, err:%v\n", err)
		config.Error(err_info)
	}
	fmt.Println(v)
	return v
}

// 分页
func (a *Article) GetArticlesPage(pageNum int, pageSize int) ([]*Article, error) {
	var articles []*Article
	args1 := (pageNum-1) * pageSize
	err := a.db().Limit(pageSize).Offset(args1).Find(&articles).Error
	if err != nil {
		err_info :=fmt.Sprintf("Get article page failed, err:%s\n", err.Error())
		config.Error(err_info)
		panic(err_info)
	}
	return articles, nil
}

// 获取指定用户全部文章
func (a *Article) GetArticlesByPersonID(id int) ([]*Article, error) {
	var artObj Article
	var articles []*Article
	err := a.db().Where("id = ?", id).Select("person_name").Find(&artObj).Error
	if err != nil{
		err_info :=fmt.Sprintf("Get user articles failed, err:%s\n", err)
		config.Error(err_info)
		panic(err)
	}
	name := artObj.Person_name
	err = a.db().Where("person_name = ?", name).Find(&articles).Error

	return articles, nil
}

// 查找文章by title
func (a *Article) FindArticleByTitle(title string) ([]*Article, error) {
	var articles []*Article
	args := "%"+title+"%"
	err := a.db().Where("title LIKE ?",args).Find(&articles).Error
	if err != nil{
		err_info :=fmt.Sprintf("find the %s article failed, err:%s\n", title, err)
		config.Error(err_info)
		panic(err)
	}
	return articles, nil
}

// 添加新文章
func (a *Article) InsertNewArticle() {
	err := a.db().Create(&a).Error
	if err != nil{
		err_info :=fmt.Sprintf("create new article failed, err:%s\n",err)
		config.Error(err_info)
		panic(err)
	}
}

// 修改文章
func (a *Article) EditArticle() {
	err := a.db().Model(&a).Updates(&a).Error
	if err != nil{
		err_info :=fmt.Sprintf("edit article failed, err:%s\n",err)
		config.Error(err_info)
		panic(err)
	}
}

func (a *Article) DelArticle() {

	err := a.db().Delete(&a).Error
	if err != nil{
		err_info :=fmt.Sprintf("create new article failed, err:%s\n",err)
		config.Error(err_info)
		panic(err)
	}
}