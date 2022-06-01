/**
* @Author: 18209
* @Description:
* @File:  pkg
* @Version: 1.0.0
* @Date: 2022/5/29 8:58
 */

package handlers

import (
	"errors"
	"gateway/pkg/info"
	"gateway/pkg/logging"
	"gateway/services"
)

func PaincIfVedioError(err error) {
	if err != nil {
		err = errors.New("vedioService--" + err.Error())
		logging.Info(err)
		panic(err)
	}
}

//包装响应结构体
func PackageRep(respRpc *services.VedioStreamResp, resp *info.VedioStreamRespModel) {
	resp.NextTime = respRpc.NextTime
	resp.StatusCode = respRpc.StatusCode
	resp.StatusMsg = respRpc.StatusMsg
	resp.VideoList = make([]*info.Video, 0)
	for i := 0; i < len(respRpc.VideoList); i++ {
		vedio := &info.Video{
			Author: info.User{
				FollowCount:   respRpc.VideoList[i].Author.FollowCount,
				FollowerCount: respRpc.VideoList[i].Author.FollowerCount,
				UID:           respRpc.VideoList[i].Author.UID,
				IsFollow:      respRpc.VideoList[i].Author.IsFollow,
				Name:          respRpc.VideoList[i].Author.Name,
			},
			CommentCount:  respRpc.VideoList[i].CommentCount,
			CoverURL:      respRpc.VideoList[i].CoverURL,
			FavoriteCount: respRpc.VideoList[i].FavoriteCount,
			VID:           respRpc.VideoList[i].VID,
			IsFavorite:    respRpc.VideoList[i].IsFavorite,
			PlayURL:       respRpc.VideoList[i].PlayURL,
			Title:         respRpc.VideoList[i].Title,
		}
		resp.VideoList = append(resp.VideoList, vedio)
	}
}
