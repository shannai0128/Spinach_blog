package controller

import (
	"github.com/c479096292/Spinach_blog/service"
	"github.com/gin-gonic/gin"
)


func GetArticleTotal() gin.HandlerFunc {
	return func(c *gin.Context) {
		articleObj := service.Article{Name:"article"}
		c.JSON(200,articleObj.Count())
		//return articleObj.Count()
	}
}

//func GetArticlesByPersonID(c *gin.Context)  {
//	var app app.App
//	// TODO
//	c.Param("token")
//}


