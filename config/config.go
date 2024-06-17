package config

import "github.com/spf13/viper"

var Config Conf

type Conf struct {
	Minio  `mapstructure:"minio"`
	Milvus `mapstructure:"milvus"`
	Mysql  `mapstructure:"mysql"`
}
type Minio struct {
	BucketName string `mapstructure:"bucketName"`
	Endpoint   string `mapstructure:"endpoint"`
}

type Milvus struct {
	Endpoint string `mapstructure:"endpoint"`
}

type Mysql struct {
	UserName string `mapstructure:"userName"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbName"`
	Endpoint string `mapstructure:"endpoint"`
}

func InitConfig() {
	//初始化配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&Config); err != nil {
		panic(err)
	}
}
