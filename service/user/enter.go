package user

import "ant-admin/gin-angular-admin/repository"

type ServiceGroup struct {
	UserService
}

var (
	userRepository = repository.RespositoryGroupApp.UserRepositoryGroup.UserRepository
)
