package user

import (
	"ant-admin/gin-angular-admin/model/dto/sysDto"
	"ant-admin/gin-angular-admin/response"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

// Register 用户注册
func (c *UserController) Register(ctx *gin.Context) {
	var request sysDto.UserDto

	if err := ctx.ShouldBind(&request); err != nil {
		response.Fail(ctx, "数据验证错误，名称比天", nil)
		return
	}
	userService.Register(request)
}
