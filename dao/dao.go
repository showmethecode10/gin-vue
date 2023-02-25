package dao

import (
	"fmt"
	"gin-vue/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Manager interface {
	Register(user *model.User)
}

//封装数据库的db，将db变成manager的database，对数据库操作的db换成了对manager.database的操作
type manager struct {
	database *gorm.DB
}

// Mgr 提供外部调用Manager接口的变量
var Mgr Manager

func init() {
	fmt.Println("初始化数据库...")
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:@tcp(127.0.0.1:3306)/gin-vue?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	//实例化manager，将Mgr关联到数据库
	Mgr = &manager{
		database: db,
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("自动创建User表失败：", err)
	}
}
