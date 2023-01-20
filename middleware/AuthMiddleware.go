package middleware

import (
	"ant-admin/gin-angular-admin/common"
	"ant-admin/gin-angular-admin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 authorization header
		tokenString := ctx.GetHeader("Authorization")

		// 验证格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort() // 将请求抛弃
			return
		}

		// Bearer 加上空格一共7位，从7位开始截取
		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)

		// 如果解析失败，解析后的token无效
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort() // 将请求抛弃
			return
		}

		//验证通过后获取claim中的uerid
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 用户不存在则token无效
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort() // 将请求抛弃
			return
		}

		//用户存在，讲user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
