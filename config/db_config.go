package config

import (
	"os"
	"strconv"
)

type DBConnect struct {
	DBDriver string
	DBName string
	DBUser string
	DBPasswd string
	PREFIX string
}

type RedisConnect struct {
	Addr string
	Password string
	DB int
}

func InitDBConnect() *DBConnect {
	var db DBConnect
	db.DBDriver = os.Getenv("DBDRIVER")
	db.DBName = os.Getenv("DBNAME")
	db.DBUser = os.Getenv("DBUSER")
	db.DBPasswd = os.Getenv("DBPASSWD")
	db.PREFIX = "blog_"
	return &db
}

func InitRedisConnect() *RedisConnect {
	var redis RedisConnect
	redis.Addr = os.Getenv("REDIS_ADDR")
	redis.Password = os.Getenv("REDIS_PASSWORD")
	REDIS_DB := os.Getenv("REDIS_DB")
	db, err:= strconv.Atoi(REDIS_DB)
	if err!=nil{
		//log.Infof()"cover redis_db exception: %s"err)
	}
	redis.DB = db
	return &redis
}




