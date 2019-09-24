package common

import (
	"github.com/c479096292/Spinach_blog/config"
	"github.com/gin-gonic/gin"
)

type Err struct {
	ErrCode int `json:"err_code"`
	Msg string `json:"msg"`
}

//func (e Err) CustomError(err error) string {
//	return fmt.Sprintf("Error code:%s, error:%s", e.ErrCode, err.Error())
//}

type Response struct {
	c *gin.Context
	Err
	Data interface{} `json:"data"`
}

func (r Response) ApiResponse(httpCode, errCode int, data interface{}) {
	r.c.JSON(httpCode, Response{
		Err:Err{ErrCode:errCode,Msg:GetMsg(errCode)},
		Data:data,
	})
}

func (r Response) ApiFailedResponse(httpCode, errCode int, err error) {
	config.Error(err.Error())
	r.c.JSON(httpCode, Response{
		Err:Err{ErrCode:errCode,Msg:GetMsg(errCode)},
	})
}

