package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Demo2 struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	UserId    uint      `json:"user_id" gorm:"not null"`
	DemoId    uint      `json:"demoId" gorm:"not null"` // 这里gorm默认以关系名+id来做为外建
	Demo      *Demo     //Demo实体，后面可以使用关联查询    d.Db.Preload("Demo").Where("id = ?", postId).First(&post).Error
	Title     string    `json:"title" gorm:"type:varchar(50);not null"`
	HeadImg   string    `json:"head_img"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	CreateAt  time.Time `json:"create_at" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp"`
}

// 在创建之前将id赋 值
func (demo2 *Demo2) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Model(demo2).Update("ID", uuid.New())
	return
}

// 一个用户有多个信用卡，如果CreditCard，里面的UserRefer 是UserId,那么就不用手动指定外建
//type User struct{
//	gorm.Model
//	CreditCards []Creditcard `gorm:"foreignkey:UserRefer"` // 显示制定外建
//}
//
//type CreditCard struct {
//	gorm.Model
//	Number string
//	UserRefer uint
//}
