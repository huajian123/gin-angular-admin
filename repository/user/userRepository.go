package user

import (
	"ant-admin/gin-angular-admin/global"
	"ant-admin/gin-angular-admin/model/entity/sysEntity"
)

type UserRepository struct {
}

func (u UserRepository) Register(entity sysEntity.User) {
	global.GLOBAL_DB.Debug().Create(&entity)
}
