package middleware

import (
	"github.com/c479096292/Spinach_blog/common"
	"github.com/c479096292/Spinach_blog/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuth(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data utils.Claims
		code = common.SUCCESS
		 token := c.DefaultQuery("token", "")
		 if token == ""{
		 	code = common.INVALID_TOKEN
		 }else {
			claims, err:= utils.ParseToken(token)
			data = *claims
			if err != nil{
				switch err.(jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = common.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = common.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}

		 }
		 if code != common.SUCCESS{

		 	c.JSON(http.StatusUnauthorized, gin.H{
		 		"code": code,
		 		"msg": common.GetMsg(code),
		 		"data": data.Username,
			})
		 	c.Request.URL.Path = "/login"
		 	r.HandleContext(c)
			 return
		 }
		 c.Next()
	}
}