/**
* @Author: 18209
* @Description:
* @File:  vedioService
* @Version: 1.0.0
* @Date: 2022/5/27 22:39
 */

package core

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"time"
	"vedio/conf"
	"vedio/model"
	"vedio/services"
)

//获取视频流
func (*VedioService) VedioStream(ctx context.Context, req *services.VedioStreamReq, resp *services.VedioStreamResp) error {
	//查询mysql数据库，按照时间顺序返回查询数据,一次最多30个
	var vedios []model.Vedio
	err := model.DB.Select("id,uid,play_url,cover_url,favorite_count,comment_count,is_favorite,title").
		Where("upload_time > ?", req.LastestTime).Order("upload_time desc").Find(&vedios).Error
	if err != nil {
		resp.NextTime = time.Now().Unix()
		resp.StatusMsg = fmt.Sprintf("fail: %v", err)
		resp.StatusCode = 1 //查询失败
		log.Println(err)
		return err
	}

	//还需要从用户模块查询用户id

	list := make([]*services.Vedio, 0)
	var nextTime int64
	for i := 0; i < len(vedios); i++ {
		if i == len(vedios)-1 {
			nextTime = vedios[i].UploadTime
		}
		vl := &services.Vedio{
			VID: int64(vedios[i].ID),
			Author: &services.User{
				UID:           vedios[i].UId,
				Name:          "sb",
				FollowCount:   0,
				FollowerCount: 0,
				IsFollow:      false,
			},
			PlayURL:       vedios[i].PlayURL,  //视频播放地址
			CoverURL:      vedios[i].CoverURL, //视频封面地址
			FavoriteCount: vedios[i].FavoriteCount,
			CommentCount:  vedios[i].CommentCount,
			IsFavorite:    vedios[i].IsFavorite,
			Title:         vedios[i].Title,
		}
		list = append(list, vl)
	}
	resp.NextTime = nextTime
	resp.StatusMsg = "success"
	resp.StatusCode = 0
	resp.VideoList = list
	return nil
}

//视频投稿
func (*VedioService) PublishVedio(ctx context.Context, req *services.PublishReq, resp *services.PublishResp) error {
	//1、将数据上传到对象服务器
	//2、将连接存到mysql

	reader := bytes.NewReader(req.Data)

	//上传失败是否尝试重试？
	ok := model.UploadFile("vedio", req.FileName, reader, int64(len(req.Data)))
	if !ok {
		resp.StatusCode = 1
		resp.StatusMsg = "upload failed"
		return errors.New("upload vedio failed " + string(req.UId) + " " + string(req.Title))
	}

	//上传成功
	v := model.Vedio{
		UId:           req.UId,
		CommentCount:  0,
		CoverURL:      "",
		FavoriteCount: 0,
		IsFavorite:    false,
		PlayURL:       conf.ConfigYaml.Minio.Endpoint + "/vedio/" + req.FileName,
		Title:         req.Title,
		UploadTime:    time.Now().Unix(),
	}
	err := model.DB.Create(&v).Error

	return err
}

//获取发布列表
func (*VedioService) PublishList(ctx context.Context, req *services.PublishListReq, resp *services.PublishListResp) error {
	var vedios []model.Vedio
	err := model.DB.Select("id,uid,play_url,cover_url,favorite_count,comment_count,is_favorite,title,upload_time").
		Where("uid=?", req.UId).Order("upload_time desc").Find(&vedios).Error
	if err != nil {
		resp.StatusMsg = fmt.Sprintf("fail: %v", err)
		resp.StatusCode = 1 //查询失败
		log.Println(err)
		return err
	}

	//还需要从用户模块查询用户id

	list := make([]*services.Vedio, 0)
	for i := 0; i < len(vedios); i++ {
		vl := &services.Vedio{
			VID: int64(vedios[i].ID),
			Author: &services.User{ //用户信息有待补充
				UID:           vedios[i].UId,
				Name:          "sb",
				FollowCount:   0,
				FollowerCount: 0,
				IsFollow:      false,
			},
			PlayURL:       vedios[i].PlayURL,  //视频播放地址
			CoverURL:      vedios[i].CoverURL, //视频封面地址
			FavoriteCount: vedios[i].FavoriteCount,
			CommentCount:  vedios[i].CommentCount,
			IsFavorite:    vedios[i].IsFavorite,
			Title:         vedios[i].Title,
		}
		list = append(list, vl)
	}
	resp.StatusCode = 0
	resp.StatusMsg = "get publishList success"

	return nil
}

//修改点赞数,1-点赞，2-取消点赞
func (*VedioService) FavoriteCountChange(ctx context.Context, req *services.FavoriteCountChangeReq, resp *services.FavoriteChangeResp) error {
	
	return nil
}

//修改评论数
func (*VedioService) CommentCountChange(ctx context.Context, req *services.FavoriteCountChangeReq, resp *services.FavoriteChangeResp) error {

	return nil
}
