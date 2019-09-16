package config

import "os"

type DBConnect struct {
	DBName string
	DBUser string
	DBPasswd string

}

func InitDBConnect() *DBConnect {
	var db DBConnect
	db.DBName = os.Getenv("DBNAME")
	db.DBUser = os.Getenv("DBUSER")
	db.DBPasswd = os.Getenv("DBPASSWD")
	return &db
}
