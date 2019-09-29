package db

import (
	"errors"
	"fmt"
	"github.com/c479096292/Spinach_blog/config"
	"github.com/go-redis/redis"
	"time"
)

var Redisdb *redis.Client

// 初始化连接
func InitRedis()  {
	redisObj := config.InitRedisConnect()
	// TODO change to pool
	Redisdb = redis.NewClient(&redis.Options{
		Addr:     redisObj.Addr,
		//Password: redisObj.Password,
		DB:       redisObj.DB,  // use default DB
	})
	err := errors.New("")
	_, err = Redisdb.Ping().Result()
	if err != nil {
		config.Info(err)
		return
	}

	fmt.Println("connect redis success!")
}

func Set(key string, value interface{}, tim int64) (err error) {
	dura_time:=time.Duration(tim)
	err =Redisdb.Set(key, value,dura_time).Err()
	if err != nil{
		config.Error(fmt.Sprintf("redis set key error"))
		return
	}
	return nil
}

func Get(key string) (value string, err error) {
	cmd :=Redisdb.Get(key)
	if cmd == nil{
		config.Error(fmt.Sprintf("redis get key error"))
		return
	}
	value,err =cmd.Result()
	if err != nil{
		config.Error(fmt.Sprintf("redis get key result error :%s",err))
		return
	}
	return
}

func Delete(key string) (err error) {

	affect, err := Redisdb.Del(key).Result()
	if err != nil{
		config.Error(fmt.Sprintf("delete key error :%s",err))
	}
	if affect == 0{
		config.Error(fmt.Sprintf("delete key error :%s",err))
	}
	return
}

