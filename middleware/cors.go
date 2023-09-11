package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// 跨域請求 可通過域名
		okHeaders := []string{
			"http://1.34.241.74:7351",
			"http://localhost:3001",
			"http://localhost:3299",
			"http://localhost:3000",
			"http://localhost:3002",
			"http://localhost:3008",
			"http://192.168.0.204:3001",
			"http://192.168.0.204:3002",
			"http://192.168.0.120:3000",
			"http://192.168.0.114:3000",
			"http://192.168.0.114:3001",
			"http://192.168.0.114:3002",
			"http://192.168.0.104",
			"http://192.168.0.55",
			"http://192.168.0.15",
			"http://woc.hjj.sat.com.tw",
			"http://58.211.121.91",
			"https://woc.hjj.sat.com.tw",
			"https://58.211.121.91",
		}

		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
		fmt.Println("origin", origin)

		// 客戶端域名在白名單內 可通過
		for _, header := range okHeaders {
			if origin == header {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
				fmt.Println("origin", origin)
			}
		}

		// 是否存取cookie
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		method := c.Request.Method
		// 放行OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		}
		// 處理請求
		c.Next()
	}
}
