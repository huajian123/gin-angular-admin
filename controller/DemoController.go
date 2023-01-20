package controller

import (
	"ant-admin/gin-angular-admin/model"
	"ant-admin/gin-angular-admin/repository"
	"ant-admin/gin-angular-admin/response"
	"ant-admin/gin-angular-admin/vo"
	"github.com/gin-gonic/gin"
	"strconv"
)

type IDemoController interface {
	RestController
}

type DemoController struct {
	Repository repository.DemoRepository
}

func NewDemoController() IDemoController {
	repository := repository.NewDemoRepository()
	repository.DB.AutoMigrate(model.Demo{}) // 添加自动迁移

	return DemoController{Repository: repository}
}

func (d DemoController) Create(ctx *gin.Context) {
	var requestDemo vo.CreateDemoRequest
	// var requestDemo model.Demo
	// 如果使用bind也能绑定，但是报错的话，只会以普通文本的形式返回给前端
	if err := ctx.ShouldBind(&requestDemo); err != nil {
		// response.Fail(ctx, "数据验证错误，名称比天", nil)
		panic(err)
		return
	}

	//ctx.Bind(&requestDemo)
	//
	//if requestDemo.Name == "" {
	//	response.Fail(ctx, "数据验证错误，名称比天", nil)
	//}

	demo, err := d.Repository.Create(requestDemo.Name)
	if err != nil {
		response.Fail(ctx, "创建失败", nil)
		return
	}

	// d.DB.Create(&requestDemo)

	response.Success(ctx, gin.H{"demo": demo}, "")

}

func (d DemoController) Update(ctx *gin.Context) {
	// 绑定body中的参数
	//var requestDemo model.Demo
	//ctx.Bind(&requestDemo)
	//
	//if requestDemo.Name == "" {
	//	response.Fail(ctx, "数据验证错误，名称比天", nil)
	//}

	var requestDemo vo.CreateDemoRequest
	// var requestDemo model.Demo
	if err := ctx.ShouldBind(&requestDemo); err != nil {
		response.Fail(ctx, "数据验证错误，名称比天", nil)
		return
	}

	// 获取path中的参数
	demoId, _ := strconv.Atoi(ctx.Params.ByName("id")) // 强转，将string 转换成int

	updateDemo, err := d.Repository.SelectById(demoId)

	if err != nil {
		response.Fail(ctx, "不存在", nil)
		return
	}

	//更新
	// map
	//struct
	//name value
	demo, err := d.Repository.Update(*updateDemo, requestDemo.Name)
	if err != nil {
		panic(err)
	}
	// d.DB.Model(&updateDemo).Update("name", requestDemo.Name)

	response.Success(ctx, gin.H{"demo": demo}, "修改成功")
}

func (d DemoController) Show(ctx *gin.Context) {
	// 获取path中的参数
	demoId, _ := strconv.Atoi(ctx.Params.ByName("id")) // 强转，将string 转换成int

	demo, err := d.Repository.SelectById(demoId)
	// err := d.DB.First(&demo, demoId).Error
	if err != nil {
		response.Fail(ctx, "不存在", nil)
		return
	}

	response.Success(ctx, gin.H{"demo": demo}, "")
}

func (d DemoController) Delete(ctx *gin.Context) {
	// 获取path中的参数
	demoId, _ := strconv.Atoi(ctx.Params.ByName("id")) // 强转，将string 转换成int

	// gorm从这个模型中知道要删除的是那张表
	//if err := d.DB.Delete(model.Demo{}, demoId).Error; err != nil {
	//	response.Fail(ctx, "删除失败", nil)
	//	return
	//}
	if err := d.Repository.DeleteById(demoId); err != nil {
		response.Fail(ctx, "删除失败", nil)
		return
	}

	response.Success(ctx, nil, "")
}
