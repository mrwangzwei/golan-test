package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"self-test/app/common/resp"
)

//异常接收
func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
				c.JSON(http.StatusOK, resp.FailRespone(resp.UnkonwError, "system error", []string{})) //响应
				return
			}
		}()
		c.Next()
	}
}
