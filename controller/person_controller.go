package controller

import (
	"database/sql"
	"fmt"
	"github.com/c479096292/Spinach_blog/common"
	"github.com/c479096292/Spinach_blog/config"
	"github.com/c479096292/Spinach_blog/service"
	"github.com/c479096292/Spinach_blog/utils"
	"github.com/gin-gonic/gin"
	"strconv"
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

func Register(c * gin.Context)  {
	AddNewPerson(c)
	personName := c.PostForm("person_name")
	passWord := c.PostForm("password")
	md5Password := utils.EncodeMD5(passWord)
	token := utils.GenerateToken(personName, md5Password)
	// 生成全局唯一ID
	worker, err:= utils.NewWorker(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	uniqueID := worker.GetId()
	int_uniqueID := strconv.Itoa(int(uniqueID))
	c.Set("tokenID",int_uniqueID)
	c.SetCookie(int_uniqueID,token,3600,"/","127.0.0.1",false,true)
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

func ModifyPassword(c *gin.Context)  {
	id := c.PostForm("id")
	int_id := utils.ChangeAtoi(id)
	u_int := uint(int_id)
	oldPassword := c.PostForm("old_password")
	newPassword := c.PostForm("new_password")
	person := service.Person{}
	err := person.ModifyPassword(u_int,oldPassword,newPassword)
	res_obj := common.Response{}
	if err != nil{
		config.Error(err)
		res_obj.ErrCode = common.ERROR_UPDATE_PASSWORD_FAIL
		res_obj.HttpCode = 500
		res_obj.Data = err.Error()
		res_obj.ApiFailedResponse(c)
		return
	}
	res_obj.HttpCode = common.SUCCESS
	res_obj.ApiResponse(c)
	return
}
