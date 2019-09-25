package db

import (
	"errors"
	"fmt"
	"github.com/c479096292/Spinach_blog/config"
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

	fmt.Println("connect mysql success!")
}

