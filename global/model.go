package global

import (
	"gorm.io/gorm"
)

type GAA_MODEL struct {
	ID        uint           `gorm:"primarykey"`           // 主键ID
	CreatedAt int            `gorm:"autoUpdateTime:milli"` // 创建时间
	UpdatedAt int            `gorm:"autoUpdateTime:milli"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`       // 删除时间
}
