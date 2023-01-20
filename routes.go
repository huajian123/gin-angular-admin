package main

import (
	"ant-admin/gin-angular-admin/controller"
	"ant-admin/gin-angular-admin/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("api/auth/info", middleware.AuthMiddleware(), controller.Info) // 用中间件来保护用户信息借口

	demoRoutes := r.Group("/demo")
	demoController := controller.NewDemoController()
	demoRoutes.POST("", demoController.Create)
	demoRoutes.PUT("/:id", demoController.Update)
	demoRoutes.GET("/:id", demoController.Show)
	demoRoutes.DELETE("/:id", demoController.Delete)
	// demoRoutes.PATCH("/:id", demoController.Delete) // 修改，put表示替换，patch表示局部修改，修改一部分

	demo2Routes := r.Group("/demo2")
	demo2Routes.Use(middleware.AuthMiddleware())
	demo2Controller := controller.NewDemo2Controller()
	demo2Routes.POST("", demo2Controller.Create)
	demo2Routes.PUT("/:id", demo2Controller.Update)
	demo2Routes.GET("/:id", demo2Controller.Show)
	demo2Routes.DELETE("/:id", demo2Controller.Delete)
	demo2Routes.POST("/page/list", demo2Controller.PageList)

	return r
}
