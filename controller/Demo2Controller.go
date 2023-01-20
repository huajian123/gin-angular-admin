package controller

import (
	"ant-admin/gin-angular-admin/common"
	"ant-admin/gin-angular-admin/model"
	"ant-admin/gin-angular-admin/response"
	"ant-admin/gin-angular-admin/vo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type IDemo2Controller interface {
	RestController
	PageList(ctx *gin.Context) // 分页
}

type Demo2Controller struct {
	Db *gorm.DB
}

func (d Demo2Controller) PageList(ctx *gin.Context) {
	// 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1")) // 转换为int类型

	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10")) // 转换为int类型

	//分页
	var posts []model.Demo2
	// 根据创建时间来排序
	d.Db.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// 前端渲染分页总条数
	var total int64
	d.Db.Model(model.Demo2{}).Count(&total)

	response.Success(ctx, gin.H{"data": posts, "total": total}, "成功")
}

func NewDemo2Controller() IDemo2Controller {
	db := common.GetDB()
	db.AutoMigrate(model.Demo2{})
	return Demo2Controller{Db: db}
}

func (d Demo2Controller) Create(ctx *gin.Context) {
	var requestPost vo.CreateDemo2Request
	// 数据验证
	if err := ctx.ShouldBind(&requestPost); err != nil {
		log.Println(err.Error())
		response.Fail(ctx, "数据验证错误", nil)
		return
	}

	// 获取登录用户user,在路由中加入中间件 demo2Routes.Use(middleware.AuthMiddleware())
	// AuthMiddleware.go中保存了上下文信息 ctx.Set("user", user)
	user, _ := ctx.Get("user")

	// 创建Post,也就是教程中的"文章"
	post := model.Demo2{
		UserId:  user.(model.User).ID,
		DemoId:  requestPost.DemoId,
		Title:   requestPost.Title,
		HeadImg: requestPost.HeadImg,
		Content: requestPost.Content,
	}

	if err := d.Db.Create(&post).Error; err != nil {
		panic(err)
		return
	}
	response.Success(ctx, nil, "创建成功")

}

func (d Demo2Controller) Update(ctx *gin.Context) {
	var requestPost vo.CreateDemo2Request
	// 数据验证
	if err := ctx.ShouldBind(&requestPost); err != nil {
		log.Println(err.Error())
		response.Fail(ctx, "数据验证错误", nil)
		return
	}

	// 获取path中的id
	postId := ctx.Params.ByName("id")

	var post model.Demo2

	err := d.Db.Where("id = ?", postId).First(&post).Error
	if err != nil {
		response.Fail(ctx, "文章不存在", nil)
		return
	}

	// 判断当前用户是否为文章的作者
	// 获取登录用户user,在路由中加入中间件 demo2Routes.Use(middleware.AuthMiddleware())
	// AuthMiddleware.go中保存了上下文信息 ctx.Set("user", user)
	user, _ := ctx.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(ctx, "文章不属于您，请勿非法操作", nil)
		return
	}

	// 更新文章，不知道这里是不是正确的更新
	if err := d.Db.Model(&post).Updates(requestPost).Error; err != nil {
		response.Fail(ctx, "更新失败", nil)
		return
	}

	response.Success(ctx, gin.H{"post": post}, "更新成功")
}

func (d Demo2Controller) Show(ctx *gin.Context) {
	// 获取path中的id
	postId := ctx.Params.ByName("id")

	var post model.Demo2

	// preload是外建,查询的时候，将关联的Demo取出来
	err := d.Db.Preload("Demo").Where("id = ?", postId).First(&post).Error
	if err != nil {
		response.Fail(ctx, "文章不存在", nil)
		return
	}

	response.Success(ctx, gin.H{"post": post}, "成功")
}

func (d Demo2Controller) Delete(ctx *gin.Context) {
	// 获取path中的id
	postId := ctx.Params.ByName("id")

	var post model.Demo2

	err := d.Db.Where("id = ?", postId).First(&post).Error
	if err != nil {
		response.Fail(ctx, "文章不存在", nil)
		return
	}

	// 判断当前用户是否为文章的作者
	// 获取登录用户user,在路由中加入中间件 demo2Routes.Use(middleware.AuthMiddleware())
	// AuthMiddleware.go中保存了上下文信息 ctx.Set("user", user)
	user, _ := ctx.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(ctx, "文章不属于您，请勿非法操作", nil)
		return
	}

	d.Db.Delete(&post)

	response.Success(ctx, gin.H{"post": post}, "删除成功")

}
