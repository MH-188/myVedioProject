/**
* @Author: 18209
* @Description:
* @File:  mysqlInit
* @Version: 1.0.0
* @Date: 2022/5/27 22:21
 */

package model

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库链接单例
var DB *gorm.DB

// ConnectMysql 在中间件中初始化mysql链接
func ConnectMysql(connString string) {
	db, err := gorm.Open("mysql", connString)
	// Error
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	//默认不加复数
	db.SingularTable(true)
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(20)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration() //数据库迁移
}
