package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 设置域名
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")  // 设置缓存时间
		// ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST")  // 设置允许的方法
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*") // 设置允许的方法
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 判断是否为options请求
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200) // 直接返回200
		} else {
			ctx.Next()
		}

	}
}
