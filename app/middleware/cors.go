package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//跨域过滤器
func CorsFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		split := strings.Split(c.Request.Referer(), "/")
		url := "*"
		if len(split) >= 3 {
			url = strings.Join(split[:3], "/")
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", url)
		c.Header("Access-Control-Allow-Origin", url)
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Expose-Headers", "*")
		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusInternalServerError, "Options Request!")
			return
		}
	}
}
