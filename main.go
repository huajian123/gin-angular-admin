package main

import (
	"ant-admin/gin-angular-admin/initialize"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	InitConfig()

	// 初始化数据库
	initialize.InitDb()

	r := gin.Default()

	r = CollectRoute(r)

	port := viper.GetString("server.host")

	if port != "" {
		panic(r.Run(":" + port))
	}

	panic(r.Run())
}

func InitConfig() {
	// 获取当前的工作目录
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
