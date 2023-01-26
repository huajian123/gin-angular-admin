package service

import "ant-admin/gin-angular-admin/service/user"

type ServiceGroup struct {
	UserServiceGroup user.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
