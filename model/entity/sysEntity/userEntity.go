package sysEntity

import "ant-admin/gin-angular-admin/global"

type User struct {
	global.GAA_MODEL
	Name      string `json:"name" gorm:"varchar(20);not null"`
	Telephone string `json:"telephone" gorm:"varchar(11);not null"`
	Email     string `json:"email" gorm:"varchar(255)"`
	Password  string `json:"password" gorm:"varchar(255);comment:密码"`
}
