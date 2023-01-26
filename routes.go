package main

import (
	"ant-admin/gin-angular-admin/middleware"
	"ant-admin/gin-angular-admin/router"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.RecoveryMiddleware())
	apiRoot := r.Group("api/v1")
	{
		userRoute := router.RouteGroupApp.User
		userRoute.InitUserRoute(apiRoot)
	}

	return r
}
