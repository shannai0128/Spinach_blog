package middleware

import (
	"fmt"
	"github.com/c479096292/Spinach_blog/common"
	"github.com/c479096292/Spinach_blog/config"
	"github.com/c479096292/Spinach_blog/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = common.SUCCESS

		uniqueID, ok := c.Get("tokenID")
		if !ok {
			config.Warn("tokenID not exist")
		}
		strID, ok := uniqueID.(string)
		if !ok {
			config.Warn("the value of tokenID is not of type string")
		}

		token, err := c.Cookie(strID)
		if err != nil{
			config.Warn(fmt.Sprintf("the value of tokenID:%s is not exist ", strID))
		}
		if token == "" {
			code = common.INVALID_TOKEN
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				switch err.(jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = common.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = common.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
			fmt.Println("claims.username: ", claims.Username)
		}
		if code != common.SUCCESS {
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}
