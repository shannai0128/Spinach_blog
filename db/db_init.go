package db

import (
	"errors"
	"fmt"
	"github.com/c479096292/Spinach_blog/config"
	"github.com/c479096292/Spinach_blog/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB()  {
	dbobj := config.InitDBConnect()
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s",dbobj.DBUser,dbobj.DBPasswd,dbobj.DBName)
	// 也可以使用MustConnect连接不成功就panic
	err := errors.New("")
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		config.Debug(err)
		return
	}
	DB.SetMaxOpenConns(20)
	DB.SetConnMaxLifetime(60)
	defer DB.Close()
	fmt.Println("connect mysql success!")
}


func queryOneInfo(tableName,fieldName string,id int)  {
	sql := fmt.Sprintf("select * from ? where ?=?")
	// TODO,根据tablename确定
	var p model.Person
	err := DB.Get(&p, sql, tableName,fieldName,id)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s\n", p.ID, p.PersonName)
}