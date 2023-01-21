package v1

import "ant-admin/gin-angular-admin/controller/v1/user"

type ApiGroup struct {
	UserControllerGroup user.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
