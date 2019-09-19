package db

import (
	"fmt"
	"github.com/c479096292/Spinach_blog/app"
	"github.com/c479096292/Spinach_blog/config"
	"github.com/c479096292/Spinach_blog/model"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDB(app *app.App) (err error) {
	dbobj := config.InitDBConnect()
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s",dbobj.DBUser,dbobj.DBPasswd,dbobj.DBName)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60)
	return
}


func queryOneInfo(tableName,fieldName string,id int)  {
	sql := fmt.Sprintf("select * from ? where ?=?")
	// TODO,根据tablename确定
	var p model.Person
	err := db.Get(&p, sql, tableName,fieldName,id)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", p.ID, p.Name, p.Age)
}