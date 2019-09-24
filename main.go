package main

import (
	"github.com/c479096292/Spinach_blog/config"
	"github.com/c479096292/Spinach_blog/db"
	"github.com/c479096292/Spinach_blog/middleware"
	"github.com/c479096292/Spinach_blog/utils"
	"github.com/gin-gonic/gin"
)


func main() {
	file := config.ParseConfigInfo(2)

	go utils.CheckLogSize(file) // 日志监控

	db.InitDB()
	db.InitRedis()


	router :=gin.Default()
	router.Use(middleware.JwtAuth(router))
	router.Use(middleware.LimitIP())
	//v1Group :=router.Group("/v1", v1Index) // 所有v1开头的交给v1Group处理。  v1Index也可不写
	//{
	//	// 匹配路由为 /v1/home ，匹配成功后会先调用v1Index，再调用v1Sku，所以可以在v1Index里写前置逻辑
	//	v1Group.GET("/home", v1Home)
	//	v1Group.GET("/sku", v1Sku)
	//}

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200,"demo index")
	})
	router.Run(":9000")

}
