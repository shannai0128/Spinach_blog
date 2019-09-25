package middleware

import (
	"github.com/c479096292/Spinach_blog/common"
	"github.com/c479096292/Spinach_blog/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = common.SUCCESS
		 token := c.DefaultQuery("token", "")
		 if token == ""{
		 	code = common.INVALID_TOKEN
		 }else {
			_, err:= utils.ParseToken(token)
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
		 	c.Redirect(http.StatusMovedPermanently,"/login")
		 	c.Abort()
			return
		 }
		 c.Next()
	}
}