package handlers

import (
	"context"
	"fmt"
	"gateway/pkg/info"
	"gateway/pkg/utils"
	"gateway/services"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func GetVedioStream(ginCtx *gin.Context) {
	fmt.Println("handler")
	var vedioStreamReq services.VedioStreamReq
	PaincIfVedioError(ginCtx.Bind(&vedioStreamReq))

	//取出服务实例
	vedioService := ginCtx.Keys["vedioService"].(services.VedioService)

	//通过token拿到当前访问用户的id（如果处于登录状态）
	token := ginCtx.DefaultQuery("token", "")
	if token != "" {
		claims, _ := utils.ParseToken(token)
		vedioStreamReq.UId = int64(claims.Id)
	}

	//latest_time
	latestTime := ginCtx.DefaultQuery("latest_time", "")
	parseInt, err := strconv.ParseInt(latestTime, 10, 64)
	if err != nil {
		parseInt = time.Now().Unix()
	}
	vedioStreamReq.LastestTime = parseInt

	//调用服务端的函数
	vedioStreamResp, err := vedioService.VedioStream(context.Background(), &vedioStreamReq)
	if err != nil {
		PaincIfVedioError(err)
	}
	var vedioStreamRespModel info.VedioStreamRespModel
	PackageRep(vedioStreamResp, &vedioStreamRespModel)
	ginCtx.JSON(http.StatusOK, vedioStreamRespModel)
}

func PublishVedio(ginCtx *gin.Context) {
	//取出服务实例
	vedioService := ginCtx.Keys["vedioService"].(services.VedioService)

	//获得token，解析出用户id
	claims, err := utils.ParseToken(ginCtx.PostForm("token"))
	if err != nil {
		ginCtx.JSON(http.StatusOK, gin.H{
			"StatusCode": 1,
			"StatusMsg":  "User doesn't exist",
		})
		return
	}

	//获取文件
	vData, err := ginCtx.FormFile("data")
	if err != nil {
		ginCtx.JSON(http.StatusOK, gin.H{
			"StatusCode": 1,
			"StatusMsg":  err.Error(),
		})
		return
	}

	//获取视频标题
	title := ginCtx.PostForm("title")

	//得到file文件对象
	vedio, err := vData.Open()
	if err != nil {
		ginCtx.JSON(http.StatusOK, gin.H{
			"StatusCode": 1,
			"StatusMsg":  err.Error(),
		})
		return
	}
	//读取视频数据
	data, err := ioutil.ReadAll(vedio)
	if err != nil {
		ginCtx.JSON(http.StatusOK, gin.H{
			"StatusCode": 1,
			"StatusMsg":  err.Error(),
		})
		return
	}

	publishReq := services.PublishReq{
		Data:     data,
		UId:      int64(claims.Id),
		Title:    title,
		FileName: vData.Filename,
	}
	//调用视频服务grpc接口
	publishResp, err := vedioService.PublishVedio(context.Background(), &publishReq)
	if err != nil {
		ginCtx.JSON(http.StatusOK, gin.H{
			"StatusCode": 1,
			"StatusMsg":  err.Error(),
		})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{
		"StatusCode": publishResp.StatusCode,
		"StatusMsg":  publishResp.StatusMsg,
	})
}

func PublishList(ginCtx *gin.Context) {
	//取出服务实例
	vedioService := ginCtx.Keys["vedioService"].(services.VedioService)

	//获得token，解析出请求用户id
	_, err := utils.ParseToken(ginCtx.DefaultQuery("token", ""))
	if err != nil {
		ginCtx.JSON(http.StatusOK, gin.H{
			"StatusCode": 1,
			"StatusMsg":  "Login information invalid",
			"video_list": nil,
		})
		return
	}
	userId := ginCtx.Query("user_id")
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		ginCtx.JSON(http.StatusOK, gin.H{
			"StatusCode": 1,
			"StatusMsg":  "user_id invalid",
			"video_list": nil,
		})
		return
	}
	req := &services.PublishListReq{
		UId: uid,
	}
	list, err := vedioService.PublishList(context.Background(), req)
	if err != nil {
		ginCtx.JSON(http.StatusOK, gin.H{
			"StatusCode": list.StatusCode,
			"StatusMsg":  list.StatusMsg,
			"video_list": nil,
		})
		return
	}
	ginCtx.JSON(http.StatusOK, gin.H{
		"StatusCode": list.StatusCode,
		"StatusMsg":  list.StatusMsg,
		"video_list": list.VideoList,
	})

}
