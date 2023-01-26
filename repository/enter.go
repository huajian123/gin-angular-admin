package repository

import "ant-admin/gin-angular-admin/repository/user"

type RespositoryGroup struct {
	UserRepositoryGroup user.RepositoryGroup
}

var RespositoryGroupApp = new(RespositoryGroup)
