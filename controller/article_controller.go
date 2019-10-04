package controller

import (
	"github.com/c479096292/Spinach_blog/common"
	"github.com/c479096292/Spinach_blog/config"
	"github.com/c479096292/Spinach_blog/service"
	"github.com/c479096292/Spinach_blog/utils"
	"github.com/gin-gonic/gin"
)


func GetArticleTotal(c *gin.Context) {
		article := service.Article{}
		articles_count := article.Count()
		res_obj := common.Response{HttpCode:common.SUCCESS,Data:articles_count}
		res_obj.ApiResponse(c)
}

func GetArticlePaged() gin.HandlerFunc {
	return func(c *gin.Context) {
		pageNum := c.PostForm("pageNum")
		pageSize := c.PostForm("pageSize")
		int_pageNum := utils.ChangeAtoi(pageNum)
		int_pageSize := utils.ChangeAtoi(pageSize)
		article := service.Article{}
		article_data := article.ArticlesPaged(int_pageNum,int_pageSize)
		res_obj := common.Response{HttpCode:common.SUCCESS,Data:article_data}
		res_obj.ApiResponse(c)
	}
}

func GetArticlesByPersonID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.PostForm("id")
		int_id := utils.ChangeAtoi(id)
		article := service.Article{}
		article_data := article.GetArticlesByPersonID(int_id)
		res_obj := common.Response{HttpCode:common.SUCCESS,Data:article_data}
		res_obj.ApiResponse(c)
	}
}

func FindArticleByTitle() gin.HandlerFunc {
	return func(c *gin.Context) {
		title := c.PostForm("title")
		articleObj := service.Article{}
		result_obj := articleObj.FindArticleByTitle(title)
		_, ok := result_obj.(string)
		if ok {
			res_obj := common.ErrParams
			res_obj.ApiFailedResponse(c)
			return
		}
		res_obj := common.Response{HttpCode:common.SUCCESS, Data:result_obj}
		res_obj.ApiResponse(c)
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

		article := service.Article{
			Category_id: uint_Category_id,
			Content: Content,
			Title: Title,
			View_count: int_View_count,
			Person_name: Person_name,
			Summary: Summary,
			Origin: int_Origin,
		}
		err := article.InsertNewArticle()
		res_obj := common.Response{}
		if err != nil{
			config.Error(err)
			res_obj.ErrCode = common.ERROR_ADD_ARTICLE_FAIL
			res_obj.ApiFailedResponse(c)
			return
		}
		res_obj.HttpCode = common.SUCCESS
		res_obj.ApiResponse(c)
		return
	}
}

// 修改文章
func EditArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		aid := c.PostForm("aid")
		int_aid := utils.ChangeAtoi(aid)
		uint_aid := uint(int_aid)

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

		article := service.Article{
			ID: uint_aid,
			Category_id: uint_Category_id,
			Content: Content,
			Title: Title,
			View_count: int_View_count,
			Person_name: Person_name,
			Summary: Summary,
			Origin: int_Origin,
		}
		err := article.EditArticle()
		res_obj := common.Response{}
		if err != nil{
			config.Error(err)
			res_obj.ErrCode = common.ERROR_EDIT_ARTICLE_FAIL
			res_obj.ApiFailedResponse(c)
			return
		}
		res_obj.HttpCode = common.SUCCESS
		res_obj.ApiResponse(c)
		return
	}
}

func DelArticle(c *gin.Context)  {
	aid := c.PostForm("aid")
	int_aid := utils.ChangeAtoi(aid)
	u_id := uint(int_aid)
	article := service.Article{}
	err := article.DelArticle(u_id)
	res_obj := common.Response{}
	if err != nil{
		config.Error(err)
		res_obj.ErrCode = common.ERROR_DELETE_ARTICLE_FAIL
		res_obj.ApiFailedResponse(c)
		return
	}
	res_obj.HttpCode = common.SUCCESS
	res_obj.ApiResponse(c)
	return
}