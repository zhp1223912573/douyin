package config

import (
	"testing"
	"log"
)

//测试默认配置是否正常读取
func TestDefaultConfig(t *testing.T){
	
	var url = DefaultConfig.Get("datasource.url")
	var driver = DefaultConfig.Get("datasource.driver")
	var password = DefaultConfig.Get("datasource.password")

	
	log.Println(url,driver,password)
	
}