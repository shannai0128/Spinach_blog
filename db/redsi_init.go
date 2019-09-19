package db

import (
	"errors"
	"fmt"
	"github.com/c479096292/Spinach_blog/config"
	"github.com/go-redis/redis"
	"time"
)

var redisdb *redis.Client

// 初始化连接
func InitRedis()  {
	redisObj := config.InitRedisConnect()
	// TODO gaizao wei pool
	redisdb = redis.NewClient(&redis.Options{
		Addr:     redisObj.Addr,
		//Password: redisObj.Password,
		DB:       redisObj.DB,  // use default DB
	})
	err := errors.New("")
	_, err = redisdb.Ping().Result()
	if err != nil {
		config.Info(err)
	}
	defer redisdb.Close()
	fmt.Println("connect redis success!")
}

func Set(key string, value interface{}, tim int64) (err error) {
	dura_time:=time.Duration(tim)
	err =redisdb.Set(key, value,dura_time).Err()
	if err != nil{
		return
	}
	defer redisdb.Close()
	return nil
}

func Get(key string) (value string, err error) {
	cmd :=redisdb.Get(key)
	if cmd == nil{
		fmt.Println("key is not exist")
	}
	value,err =cmd.Result()
	defer redisdb.Close()
	return
}

func Delete(key string) (err error) {

	err = redisdb.Del(key).Err()
	if err != nil{
		fmt.Printf("delete key error :%s",err)
	}
	defer redisdb.Close()
	return
}

