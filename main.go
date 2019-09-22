package main

import (
	"github.com/c479096292/Spinach_blog/config"
	"github.com/c479096292/Spinach_blog/db"
	"github.com/c479096292/Spinach_blog/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	file := config.ParseConfigInfo(2)

	go utils.CheckLogSize(file) // 日志监控

	db.InitDB()
	db.InitRedis()
	//go utils.CheckLogSize()
	// 创建一个默认的路由引擎

	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run(":9000")


}
