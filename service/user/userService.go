package user

import (
	"ant-admin/gin-angular-admin/model/dto/sysDto"
	"ant-admin/gin-angular-admin/model/entity/sysEntity"
	"github.com/jinzhu/copier"
)

type UserService struct {
}

func (u UserService) Register(dto sysDto.UserDto) {
	var entity sysEntity.User
	err := copier.Copy(&entity, &dto)
	if err != nil {
		return
	}
	userRepository.Register(entity)
}
