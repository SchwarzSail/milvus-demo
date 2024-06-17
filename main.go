// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/spf13/viper"
	"milvus-demo/biz/rpc"
	"milvus-demo/config"
	"milvus-demo/dal"
	"milvus-demo/internal/utils"
)

func init() {
	//配置文件初始化
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&config.Config); err != nil {
		panic(err)
	}
	utils.InitLog()
	dal.Init()
	rpc.Init()
}

func main() {
	h := server.Default(
		server.WithHostPorts("localhost:8888"),
		server.WithMaxRequestBodySize(419430400),
	)
	register(h)
	h.Spin()
}
