package middleware

import (
	"fmt"
	"github.com/c479096292/Spinach_blog/config"
	"github.com/c479096292/Spinach_blog/db"
	"github.com/gin-gonic/gin"
	"strconv"
)

func LimitIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		request_ip := c.ClientIP()
		values,err := db.Get(request_ip)
		if err != nil{
			err = db.Set(request_ip,0,config.ConfObj.IpBlackExpire)
			if err != nil{
				config.Error(fmt.Sprintf("set %s key error: %s",request_ip,err))
			}
		}else {
			val,err := strconv.Atoi(values)
			if err != nil{
				config.Error(fmt.Sprintf("convert %s to int error: %s", err))
			}
			if val >= 100{
				c.JSON(500,"Please don't try too fast")
			}
			db.Redisdb.Incr(request_ip) // 请求次数+1
			c.Next()
		}
	}
}
