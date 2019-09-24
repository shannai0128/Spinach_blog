package controller

import (
	"github.com/c479096292/Spinach_blog/app"
	"github.com/c479096292/Spinach_blog/service"
	"github.com/gin-gonic/gin"
)


func GetArticleTotal() int {
	articleObj := service.Article{}
	return articleObj.Count()
}

func GetArticles(c *gin.Context)  {
	var app app.App
	// TODO
	c.Param("token")
}


