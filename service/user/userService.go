package user

import (
	"ant-admin/gin-angular-admin/model/dto/sysDto"
	"ant-admin/gin-angular-admin/model/entity/sysEntity"
)

type UserService struct {
}

func (u UserService) Register(dto sysDto.UserDto) {
	var entity sysEntity.User
	entity.Email = dto.Email
	entity.Name = dto.Name
	entity.Telephone = dto.Telephone
	entity.Password = dto.Password

	userRepository.Register(entity)
}
