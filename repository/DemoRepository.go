package repository

import (
	"ant-admin/gin-angular-admin/common"
	"ant-admin/gin-angular-admin/model"
	"gorm.io/gorm"
)

// 用来抽离通用逻辑
// 可以像control中一样定义接口
type DemoRepository struct {
	DB *gorm.DB
}

func NewDemoRepository() DemoRepository {
	return DemoRepository{DB: common.GetDB()}
}

func (d DemoRepository) Create(name string) (*model.Demo, error) {
	demo := model.Demo{
		Name: name,
	}

	if err := d.DB.Create(&demo).Error; err != nil {
		return nil, err
	}
	return &demo, nil
}

func (d DemoRepository) Update(demo model.Demo, name string) (*model.Demo, error) {
	if err := d.DB.Model(&demo).Update("name", name).Error; err != nil {
		return nil, err
	}
	return &demo, nil
}

func (d DemoRepository) SelectById(id int) (*model.Demo, error) {
	var demo model.Demo
	if err := d.DB.First(&demo, id).Error; err != nil {
		return nil, err
	}
	return &demo, nil
}

func (d DemoRepository) DeleteById(id int) error {
	if err := d.DB.Delete(model.Demo{}, id).Error; err != nil {
		return err
	}
	return nil
}
