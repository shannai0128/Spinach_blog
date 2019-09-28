package controller

import (
	"database/sql"
	"github.com/c479096292/Spinach_blog/common"
	"github.com/c479096292/Spinach_blog/config"
	"github.com/c479096292/Spinach_blog/service"
	"github.com/c479096292/Spinach_blog/utils"
	"github.com/gin-gonic/gin"
)

func GetPersonTotal(c *gin.Context) {
	person := service.Person{Name:"person"}
	person_count := person.Count()
	res_obj := common.Response{HttpCode:common.SUCCESS,Data:person_count}
	res_obj.ApiResponse(c)
}

func GetPersonPage(c *gin.Context)  {
	pageNum := c.PostForm("page_num")
	pageSize := c.PostForm("page_size")
	int_pageNum := utils.ChangeAtoi(pageNum)
	int_pageSize := utils.ChangeAtoi(pageSize)
	person := service.Person{}
	person_data := person.GetPersonPage(int_pageNum, int_pageSize)
	res_obj := common.Response{
		HttpCode:common.SUCCESS,
		Data:person_data,
	}
	res_obj.ApiResponse(c)
}

func FindPersonByID(c *gin.Context)  {
	id := c.PostForm("id")
	int_id := utils.ChangeAtoi(id)
	person := service.Person{}
	person_data := person.FindPersonByID(int_id)
	res_obj := common.Response{
		HttpCode:common.SUCCESS,
		Data:person_data,
	}
	res_obj.ApiResponse(c)
}

func AddNewPerson(c *gin.Context)  {
	// TODO 注册逻辑
	res_obj := common.Response{}
	personName := c.PostForm("person_name")
	idCard := c.PostForm("id_card")
	passWord := c.PostForm("password")
	ipAddr := c.ClientIP()
	loginIP := sql.NullString{ipAddr,true}
	gender := c.PostForm("gender")
	int_gender := utils.ChangeAtoi(gender)

	if len(passWord) < 8 {
		res_obj.ErrCode = common.ERROR_ADD_USER_FAIL
		res_obj.HttpCode = 400
		res_obj.Data = "密码不得少于8位"
		res_obj.ApiFailedResponse(c)
		return
	}
	if len(idCard) != 18{
		res_obj.ErrCode = common.ERROR_ADD_USER_FAIL
		res_obj.HttpCode = 400
		res_obj.Data = "请输入有效身份证号"
		res_obj.ApiFailedResponse(c)
		return
	}

	person := service.Person{
		PersonName:personName,
		IdCard:idCard,
		PassWord:passWord,
		LoginIP:loginIP,
		Gender:int_gender,
	}
	err := person.AddtNewPerson()

	if err != nil{
		config.Error(err)
		res_obj.ErrCode = common.ERROR_ADD_USER_FAIL
		res_obj.HttpCode = 500
		res_obj.Data = err.Error()
		res_obj.ApiFailedResponse(c)
		return
	}
	res_obj.HttpCode = common.SUCCESS
	res_obj.ApiResponse(c)
}

func DelPerson(c *gin.Context)  {
	id := c.PostForm("id")
	int_id := utils.ChangeAtoi(id)
	person := service.Person{}
	err := person.DelPersonByID(int_id)
	res_obj := common.Response{}
	if err != nil{
		config.Error(err)
		res_obj.ErrCode = common.ERROR_ADD_USER_FAIL
		res_obj.HttpCode = 500
		res_obj.Data = err.Error()
		res_obj.ApiFailedResponse(c)
		return
	}
	res_obj.HttpCode = common.SUCCESS
	res_obj.ApiResponse(c)
	return
}
