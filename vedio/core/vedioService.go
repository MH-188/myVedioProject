/**
* @Author: 18209
* @Description:
* @File:  vedioService
* @Version: 1.0.0
* @Date: 2022/5/27 22:39
 */

package core

import (
	"context"
	"vedio/services"
)

//获取视频流
func (*VedioService) VedioStream(ctx context.Context, req *services.VedioStreamReq, resp *services.VedioStreamResp) error {
	
}

//视频投稿
func (*VedioService) PublishVedio(ctx context.Context, req *services.PublishReq, resp *services.PublishResp) error {

}

//获取发布列表
func (*VedioService) PublishList(ctx context.Context, req *services.PublishListReq, resp *services.PublishListResp) error {

}

//修改点赞数
func (*VedioService) FavoriteCountChange(ctx context.Context, req *services.FavoriteCountChangeReq, resp *services.FavoriteChangeResp) error {

}

//修改评论数
func (*VedioService) CommentCountChange(ctx context.Context, req *services.FavoriteCountChangeReq, resp *services.FavoriteChangeResp) error {

}
