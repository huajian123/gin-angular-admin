package user

import (
	v1 "ant-admin/gin-angular-admin/controller/v1"
	"github.com/gin-gonic/gin"
)

type UserRoute struct{}

func (r *UserRoute) InitUserRoute(Route *gin.RouterGroup) {
	route := Route.Group("user")
	userControllers := v1.ApiGroupApp.UserControllerGroup.UserController
	{
		route.POST("register", userControllers.Register)
	}
}
