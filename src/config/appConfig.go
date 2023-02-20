package config

import "github.com/spf13/viper"

//全局配置
var DefaultConfig = viper.New()

//初始化全局配置文件
func init() {
	//配置默认读取配置路径
	DefaultConfig.AddConfigPath("../../resources")
	DefaultConfig.SetConfigName("application") 
	DefaultConfig.SetConfigType("yaml")       

	//尝试进行配置读取
	if err := DefaultConfig.ReadInConfig(); err != nil {
		panic(err)
	}

}