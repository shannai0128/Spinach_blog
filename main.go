package main

import (
	"fmt"
	"github.com/c479096292/Spinach_blog/config"
	"github.com/c479096292/Spinach_blog/db"
	"github.com/c479096292/Spinach_blog/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)


func main() {
	file := config.ParseConfigInfo(2)

	go utils.CheckLogSize(file) // 日志监控

	db.InitDB()
	db.InitRedis()


	router :=gin.Default()
	//router.Use(middleware.LimitIP())
	//v1Group :=router.Group("/v1", v1Index) // 所有v1开头的交给v1Group处理。  v1Index也可不写
	//{
	//	// 匹配路由为 /v1/home ，匹配成功后会先调用v1Index，再调用v1Sku，所以可以在v1Index里写前置逻辑
	//	v1Group.GET("/home", v1Home)
	//	v1Group.GET("/sku", v1Sku)
	//}

	router.GET("/hello", func(c *gin.Context) {
		request_ip := c.ClientIP()
		values,err := db.Get(request_ip)
		if err != nil{
			err = db.Set(request_ip,0,60000000000)
			if err != nil{
				config.Error(fmt.Sprintf("set %s key error: %s",request_ip,err))

			}
		}else {
			val,err := strconv.Atoi(values)
			if err != nil{
				config.Error(fmt.Sprintf("convert %s to int error: %s", err))
			}
			if val >= 100{
				c.JSON(500,"Please don't try too fast")
			}
			db.Redisdb.Incr(request_ip) // 请求次数+1

		}
	})
	router.Run(":9000")

}
