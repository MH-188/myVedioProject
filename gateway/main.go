/**
* @Author: 18209
* @Description:
* @File:  main
* @Version: 1.0.0
* @Date: 2022/5/29 0:11
 */

package main

import (
	"gateway/services"
	"gateway/weblib"
	"gateway/wrappers"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"time"
)

func main() {
	etcdReq := etcd.NewRegistry(
		registry.Addrs("192.168.1.106:2379"),
	)

	//视频
	vedioMicroService := micro.NewService(
		micro.Name("vedioService.Client"),
		micro.WrapClient(wrappers.NewVedioWrapper),
	)
	//视频服务调用示例
	vedioService := services.NewVedioService("rpcVedioService", vedioMicroService.Client())

	//创建微服务实例
	server := web.NewService(
		web.Name("httpService"),
		web.Address(":4000"),
		//将微服务调用实例使用gin处理
		web.Handler(weblib.NewRouter(vedioService)), //将服务传进去
		web.Registry(etcdReq),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	//接收命令行参数
	_ = server.Init()
	_ = server.Run()
}
