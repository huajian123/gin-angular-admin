package user

import "ant-admin/gin-angular-admin/service"

type ApiGroup struct {
	UserController
}

var (
	userService = service.ServiceGroupApp.UserServiceGroup.UserService
)
