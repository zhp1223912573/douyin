package dao

import(
	"testing"

	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
)


type Student struct {
	gorm.Model
	Name string
	Age  uint64
}

func TestConnectMysql(t *testing.T){
	//建表
	// Db.AutoMigrate(&Student{})
	
	//插入数据
	Db.Create(&Student{
		Name: "zhangsan",
		Age: 60,
	})

	//查询数据



}