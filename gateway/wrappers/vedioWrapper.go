/**
* @Author: 18209
* @Description:
* @File:  vedioWrapper
* @Version: 1.0.0
* @Date: 2022/5/29 11:33
 */

package wrappers

import (
	"context"
	"gateway/services"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
)

func NewVedio(id int64, name string, playUrl string) *services.Vedio {
	return &services.Vedio{
		VID:   id,
		Title: name,
		Author: &services.User{
			FollowCount:   0,
			FollowerCount: 0,
			UID:           0,
			IsFollow:      false,
			Name:          "抖音小红娘",
		},
		CommentCount:  0,
		CoverURL:      "",
		FavoriteCount: 0,
		IsFavorite:    false,
		PlayURL:       playUrl,
	}
}

//降级函数
func DefaultVedio(resp interface{}) {
	models := make([]*services.Vedio, 0)
	models = append(models, NewVedio(0, "降级视频流", "http://192.168.1.106:9000/vedio/20220518.mp4"))
	result := resp.(*services.VedioStreamResp)
	result.VideoList = models
}

type VedioWrapper struct {
	client.Client
}

func (wrapper *VedioWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()
	config := hystrix.CommandConfig{
		Timeout:                3000,
		RequestVolumeThreshold: 20,   //熔断器请求阈值，默认20，意思是有20个请求才能进行错误百分比计算
		ErrorPercentThreshold:  50,   //错误百分比，当错误超过百分比时，直接进行降级处理，直至熔断器再次 开启，默认50%
		SleepWindow:            5000, //过多长时间，熔断器再次检测是否开启，单位毫秒ms（默认5秒）
	}
	hystrix.ConfigureCommand(cmdName, config)
	return hystrix.Do(cmdName, func() error {
		return wrapper.Client.Call(ctx, req, rsp)
	}, func(err error) error {
		DefaultVedio(rsp)
		return err
	})
}

//NewProductWrapper 初始化Wrapper
func NewVedioWrapper(c client.Client) client.Client {
	return &VedioWrapper{c}
}
