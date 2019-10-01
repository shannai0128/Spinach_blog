package controller

import (
	"database/sql"
	"fmt"
	"github.com/c479096292/Spinach_blog/common"
	"github.com/c479096292/Spinach_blog/config"
	"github.com/c479096292/Spinach_blog/db"
	"github.com/c479096292/Spinach_blog/service"
	"github.com/c479096292/Spinach_blog/utils"
	"github.com/gin-gonic/gin"
	"math/rand"
)

func AddPersonInfo(c *gin.Context)  {
	res_obj := common.Response{}
	id := c.PostForm("id")
	personEmail := c.PostForm("person_email")
	personAddr := c.PostForm("person_addr")
	nickName := c.PostForm("nick_name")
	mobile := c.PostForm("mobile")

	// 验证短信验证码
	verifyCode := c.PostForm("verify_code")

	acquire_vCode, err := db.Get("verify_"+mobile)

	if acquire_vCode != verifyCode {
		config.Error("acquire verify code error: ",err)
		res_obj.ErrCode = common.ERROR_GET_VERIFY_CODE_FAIL
		res_obj.HttpCode = 400
		res_obj.Data = err.Error()
		res_obj.ApiFailedResponse(c)
		return
	}

	int_id := utils.ChangeAtoi(id)

	personInfo := service.PersonInfo{
		ID: int_id,
		Person_email: sql.NullString{personEmail,true},
		Person_addr: sql.NullString{personAddr, true},
		Nick_name: sql.NullString{nickName, true},
		Mobile: sql.NullString{mobile, true},
	}
	err = personInfo.AddtPersonDetail()

	if err != nil{
		config.Error(err)
		res_obj.ErrCode = common.ERROR_ADD_USER_INFO_FAIL
		res_obj.HttpCode = 500
		res_obj.Data = err.Error()
		res_obj.ApiFailedResponse(c)
		return
	}
	res_obj.HttpCode = common.SUCCESS
	res_obj.ApiResponse(c)
	// 验证通过后删除
	err = db.Delete("verify_"+mobile)
	if err != nil {
		config.Error(err)
		return
	}
}

func EditPersonInfo(c *gin.Context) {
	res_obj := common.Response{}
	id := c.PostForm("id")
	personEmail := c.PostForm("person_email")
	personAddr := c.PostForm("person_addr")
	nickName := c.PostForm("nick_name")
	int_id := utils.ChangeAtoi(id)
	mobile := c.PostForm("mobile")
	personInfo := service.PersonInfo{
		ID: int_id,
		Person_email: sql.NullString{personEmail,true},
		Person_addr: sql.NullString{personAddr, true},
		Nick_name: sql.NullString{nickName, true},
		Mobile: sql.NullString{mobile, true},
	}
	err := personInfo.EditPersonInfo()
	if err != nil{
		config.Error(err)
		res_obj.ErrCode = common.ERROR_EDIT_USER_FAIL
		res_obj.HttpCode = 500
		res_obj.Data = err.Error()
		res_obj.ApiFailedResponse(c)
		return
	}
	res_obj.HttpCode = common.SUCCESS
	res_obj.ApiResponse(c)
}

// 短信验证码
func VerifyCode(c *gin.Context)  {
	mobile := c.PostForm("mobile")
	verifyCode := fmt.Sprintf("%06v",rand.Int31n(1000000))
	_ = db.Set("verify_"+mobile,verifyCode,config.ConfObj.VerifyCodeExpire)
	res_obj := common.Response{}

	res_obj.HttpCode = common.SUCCESS
	res_obj.ApiResponse(c)
}

