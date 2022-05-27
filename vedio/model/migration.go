/**
* @Author: 18209
* @Description:
* @File:  migration
* @Version: 1.0.0
* @Date: 2022/5/27 22:35
 */

package model

func migration() {
	//自动迁移模式
	DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&Vedio{})
}
