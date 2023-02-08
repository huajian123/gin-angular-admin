package user

import (
	"ant-admin/gin-angular-admin/model/dto/sysDto"
	"ant-admin/gin-angular-admin/response"
	"ant-admin/gin-angular-admin/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strings"
)

type UserController struct{}

func mapToJson(param map[string]string) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

// Register 用户注册
func (c *UserController) Register(ctx *gin.Context) {
	var request sysDto.UserDto

	err := ctx.ShouldBind(&request)
	if err != nil {
		response.Fail(ctx, mapToJson(util.TransError(err)), nil)
		return
	}
	userService.Register(request)
	response.Success(ctx, nil, "注册成功")
}
