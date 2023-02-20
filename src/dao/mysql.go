package dao

import (
	. "douyin/src/config"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	
	// "gopkg.in/yaml.v2"
	// "io/ioutil"
)



var Db *gorm.DB


// DRIVER 指定驱动
// const DRIVER = "mysql"

// // 配置参数映射结构体
// type conf struct {
// 	Url      string `yaml:"url"`
// 	UserName string `yaml:"userName"`
// 	PassWord string `yaml:"passWord"`
// 	DbName   string `yaml:"dbname"`
// 	Port     string `yaml:"post"`
// }

// // 获取配置参数数据
// func (c *conf) getConf() *conf {
// 	// 读取resources/application.yaml文件
// 	yamlFile, err := ioutil.ReadFile("resources/application.yaml")
// 	// 若出现错误，打印错误提示
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	// 将读取的字符串转换成结构体conf
// 	err = yaml.Unmarshal(yamlFile, c)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	return c
// }

//通过默认配置文件初始化mysql
// // InitMySql 初始化连接数据库
// func InitMySql() (err error) {
// 	var c conf
// 	// 获取yaml配置参数
// 	conf := c.getConf()
// 	// 将yaml配置参数拼接成连接数据库的url
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 		conf.UserName,
// 		conf.PassWord,
// 		conf.Url,
// 		conf.Port,
// 		conf.DbName,
// 	)
// 	// 连接数据库
// 	SqlSession, err = gorm.Open(DRIVER, dsn)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// 验证数据库连接是否成功，若成功，则无异常
// 	return SqlSession.DB().Ping()
// }

//获取默认配置 进行数据库链接
func init(){
	//获取数据库配置
	var host = DefaultConfig.Get("datasource.url")
	var driver = DefaultConfig.Get("datasource.driver").(string)
	var userName = DefaultConfig.Get("datasource.userName")
	var password = DefaultConfig.Get("datasource.password")
	var dbName = DefaultConfig.Get("datasource.dbName")

	log.Printf("连接%s数据库,链接为：%s 密码%d",driver,host,password)

	//拼接url
	url := fmt.Sprintf("%s:%d@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		userName,
		password,
		host,
		dbName)

	//建立连接
	db,err := gorm.Open(driver,url)
	if err != nil {
		panic(err)
	}

	if db.DB().Ping() != nil {
		log.Printf("数据库连接成功！")
	}

	Db = db

}

// Close 关闭数据库连接
func Close() {
	err := Db.Close()
	if err != nil {
		return
	}
}
