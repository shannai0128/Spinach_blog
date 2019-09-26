package controller

import (
	"github.com/c479096292/Spinach_blog/config"
	"github.com/c479096292/Spinach_blog/service"
	"github.com/c479096292/Spinach_blog/utils"
	"github.com/gin-gonic/gin"
)


func GetArticleTotal() gin.HandlerFunc {
	return func(c *gin.Context) {
		articleObj := service.Article{Name:"article"}
		c.JSON(200,articleObj.Count())
		//return articleObj.Count()
	}
}

func GetArticlePaged() gin.HandlerFunc {
	return func(c *gin.Context) {
		pageNum := c.PostForm("pageNum")
		pageSize := c.PostForm("pageSize")
		int_pageNum := utils.ChangeAtoi(pageNum)
		int_pageSize := utils.ChangeAtoi(pageSize)
		articleObj := service.Article{}
		c.JSON(200, articleObj.ArticlesPaged(int_pageNum,int_pageSize))
	}
}

func GetArticlesByPersonID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.PostForm("id")
		int_id := utils.ChangeAtoi(id)

		articleObj := service.Article{}
		c.JSON(200, articleObj.GetArticlesByPersonID(int_id))
	}
}

func FindArticleByTitle() gin.HandlerFunc {
	return func(c *gin.Context) {
		title := c.PostForm("title")
		articleObj := service.Article{}
		result_obj := articleObj.FindArticleByTitle(title)
		result, ok := result_obj.(string)
		if ok {
			c.JSON(200, result)
			return
		}
		c.JSON(200, result_obj)
		return
	}
}

func CreateNewArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		Category_id := c.PostForm("category_id")
		int_Category_id := utils.ChangeAtoi(Category_id)
		uint_Category_id := uint(int_Category_id)

		Content := c.PostForm("content")
		Title := c.PostForm("title")
		View_count := c.PostForm("view_count")
		int_View_count := utils.ChangeAtoi(View_count)
		Person_name := c.PostForm("person_name")
		Summary := c.PostForm("summary")
		Origin := c.PostForm("origin")
		int_Origin := utils.ChangeAtoi(Origin)
		Praise := c.PostForm("praise")
		int_Praise := utils.ChangeAtoi(Praise)

		article := service.Article{
			Category_id: uint_Category_id,
			Content: Content,
			Title: Title,
			View_count: int_View_count,
			Person_name: Person_name,
			Summary: Summary,
			Origin: int_Origin,
			Praise: int_Praise,
		}
		err := article.InsertNewArticle()
		if err != nil{
			config.Error(err)
			return
		}
		return
	}
}