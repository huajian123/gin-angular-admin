package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name      string    `gorm:"type:varchar(20);not null"`
	Telephone string    `gorm:"varchar(11);not null;unique"`
	Email     string    `gorm:"varchar(255);unique"`
	Password  string    `gorm:"size:255;not null"`
	CreateAt  time.Time `json:"create_at" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp"`
}
