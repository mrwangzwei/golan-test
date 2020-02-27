package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"self-test/app/utils"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)



				c.JSON(http.StatusOK, utils.Respone(utils.UNKNOW_ERROR, "system error", []string{}))		//响应
				return
			}
		}()
		c.Next()
	}
}
