package common

import (
	"github.com/gin-gonic/gin"
)


type Response struct {
	HttpCode int
	ErrCode int `json:"err_code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func (r *Response) ApiResponse(c *gin.Context) {
	c.JSON(r.HttpCode, Response{
		HttpCode:r.HttpCode,
		Data: r.Data,
	})
}

func (r *Response) ApiFailedResponse(c *gin.Context) {
	c.JSON(r.HttpCode, Response{
		ErrCode: r.ErrCode,
		Msg: GetMsg(r.ErrCode),
		Data: r.Data,
	})
}

func newApiError(errCode int, data interface{}) *Response {
	return &Response{
		HttpCode:errCode,
		ErrCode:errCode,
		Msg:GetMsg(errCode),
		Data:data,
	}
}

var (
	ErrDatabase = newApiError(500,"Database error, an Unknown error occurred")
	ErrParams = newApiError(400, "params error, please check and try again")
)