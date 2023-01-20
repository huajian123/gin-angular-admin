package middleware

import (
	"ant-admin/gin-angular-admin/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 这个中间件相当于将调用栈的信息返回给前端，将错误信息直接返回给前端
func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Fail(ctx, fmt.Sprint(err), nil)
			}
		}()

		ctx.Next()
	}
}
