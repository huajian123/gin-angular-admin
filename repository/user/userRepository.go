package user

import (
	"ant-admin/gin-angular-admin/global"
	"ant-admin/gin-angular-admin/model/entity/sysEntity"
)

type UserRepository struct {
}

func (u UserRepository) Register(entity sysEntity.User) {
	err := global.GLOBAL_DB.Debug().Create(&entity).Error

	// 在RecoveryMiddleware这个中间件中，拦截这个错误，如果报错，将会吧这个错误在前端展示出来
	if err != nil {
		panic(err)
		return
	}
}
