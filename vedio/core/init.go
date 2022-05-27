/**
* @Author: 18209
* @Description:
* @File:  init
* @Version: 1.0.0
* @Date: 2022/5/27 22:32
 */

package core

import (
	"vedio/conf"
	"vedio/model"
)

func Init() {
	//加载配置文件
	err := conf.LoadConfig()
	if err != nil {
		panic(err)
	}
	model.ConnectMysql(conf.ConfigYaml.Mysql.DataSource)
}
