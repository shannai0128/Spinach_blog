package main

import (
	"github.com/c479096292/Spinach_blog/config"
	"github.com/c479096292/Spinach_blog/controller"
	"github.com/c479096292/Spinach_blog/db"
	"github.com/c479096292/Spinach_blog/middleware"
	"github.com/c479096292/Spinach_blog/utils"
	"github.com/gin-gonic/gin"
)

func lst(c *gin.Context)  {
	c.JSON(200, "dasdsa")
}

func main() {
	file := config.ParseConfigInfo(2)

	go utils.CheckLogSize(file) // 日志监控

	db.InitDB()
	db.InitRedis()


	router :=gin.Default()
	router.GET("/login", func(c *gin.Context) {
		c.JSON(200,"login index")
	})
	//router.Use(middleware.JwtAuth())
	router.Use(middleware.LimitIP())

	atcGroup := router.Group("/article")
	{
		atcGroup.GET("/total", controller.GetArticleTotal)
		atcGroup.POST("/paged", controller.GetArticlePaged())
		atcGroup.POST("/articles", controller.GetArticlesByPersonID())
		atcGroup.POST("/find", controller.FindArticleByTitle())
		atcGroup.POST("/new", controller.CreateNewArticle())
		atcGroup.POST("/edit", controller.EditArticle())
		atcGroup.POST("/delete", controller.DelArticle)
	}
	userGroup := router.Group("/user")
	{
		userGroup.GET("/total",controller.GetPersonTotal)
		userGroup.POST("/new", controller.AddNewPerson)
		userGroup.POST("/paged", controller.GetPersonPage)
		userGroup.POST("/find", controller.FindPersonByID)
		userGroup.POST("/del", controller.DelPerson)
	}
	detailGroup := router.Group("/detail")
	{
		detailGroup.GET("/total",controller.GetPersonTotal)
		detailGroup.POST("/new", controller.AddNewPerson)
		detailGroup.POST("/paged", controller.GetPersonPage)
	}
	router.Run(":9000")

}
