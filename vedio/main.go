/**
* @Author: 18209
* @Description:
* @File:  main
* @Version: 1.0.0
* @Date: 2022/5/27 15:00
 */

package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"vedio/core"
	"vedio/services"
)

func main() {
	core.Init()

	//etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs("192.168.1.106:2379"),
		//registry.Addrs("192.168.24.226:2379"),
	)

	//得到一个微服务示例
	microService := micro.NewService(
		micro.Name("rpcVedioService"), //微服务名字
		micro.Address("127.0.0.1:8083"),
		micro.Registry(etcdReg), //etcd注册件
	)

	//结构命令行参数，初始化
	microService.Init()

	//服务注册
	_ = services.RegisterVedioServiceHandler(microService.Server(), new(core.VedioService))
	//启动微服务
	_ = microService.Run()
}
