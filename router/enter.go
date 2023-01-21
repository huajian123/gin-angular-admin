package router

import "ant-admin/gin-angular-admin/router/user"

type RouteGroup struct {
	User user.UserRoute
}

var RouteGroupApp = new(RouteGroup)
