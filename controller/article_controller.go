package controller

import "github.com/c479096292/Spinach_blog/service"


func GetArticleTotal() int {
	articleObj := service.Article{}
	return articleObj.Count()
}


