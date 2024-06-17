package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"milvus-demo/config"
)

var DB *gorm.DB

func InitMySQL() {
	conf := config.Config.Mysql
	dsn := conf.UserName + ":" + conf.Password + "@tcp(" + conf.Endpoint + ")/" + conf.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		Logger:         logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(&Image{})
	DB = db
}
