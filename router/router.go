package router

import (
	"github.com/c479096292/Spinach_blog/controller"
	"github.com/c479096292/Spinach_blog/middleware"
	"github.com/gin-gonic/gin"
)

func LoadRouterHandler()  {
	router :=gin.Default()
	router.Use(middleware.LimitIP())
	router.GET("/login", func(c *gin.Context) {
		c.JSON(200,"login page")
	})
	router.POST("/register", controller.Register)
	//router.Use(middleware.JwtAuth())

	router.GET("/index", func(c *gin.Context) {
		c.JSON(200,"index page")
	})

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
		userGroup.POST("/pwd", controller.ModifyPassword)
	}
	detailGroup := router.Group("/detail")
	{
		detailGroup.POST("/new", controller.AddPersonInfo)
		detailGroup.POST("/edit", controller.EditPersonInfo)
		detailGroup.POST("/vc", controller.VerifyCode)
	}
	adminGroup := router.Group("/admin")
	{
		adminGroup.POST("/add", controller.AddNewPerson)
	}

	router.Run(":9000")
}



