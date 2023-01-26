package initialize

import (
	"ant-admin/gin-angular-admin/global"
	"ant-admin/gin-angular-admin/model/entity/sysEntity"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"net/url"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc") // 时区
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc),
	)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, // 生成的表名为单数，不然会自动添加s例如users
	})
	if err != nil {
		panic("failed to connect database,err: " + err.Error())
	}

	db.AutoMigrate(&sysEntity.User{})

	DB = db
	global.GLOBAL_DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
