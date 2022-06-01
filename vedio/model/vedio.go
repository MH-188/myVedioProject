/**
* @Author: 18209
* @Description:
* @File:  vedio
* @Version: 1.0.0
* @Date: 2022/5/27 22:11
 */

package model

import "github.com/jinzhu/gorm"

//数据库模型
type Vedio struct {
	gorm.Model
	UId           int64  `gorm:"not null"`
	CommentCount  int64  `gorm:"default:'0'"`
	CoverURL      string `gorm:"not null"`
	FavoriteCount int64  `gorm:"default:'0'"`
	IsFavorite    bool   `gorm:"default:false"`
	PlayURL       string `gorm:"not null"`
	Title         string `gorm:"not null"`
	UploadTime    int64
}
