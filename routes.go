package main

import (
	"ant-admin/gin-angular-admin/router"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	apiRoot := r.Group("api/v1")
	{
		userRoute := router.RouteGroupApp.User
		userRoute.InitUserRoute(apiRoot)
	}

	return r
}
