/**
* @Author: 18209
* @Description:
* @File:  init
* @Version: 1.0.0
* @Date: 2022/5/29 9:05
 */

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//接收微服务示例，并存到gin.Key中,main函数中的实例
func InitMiddleware(service []interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		//将实例存在gin.Key中
		context.Keys = make(map[string]interface{})
		//context.Keys["userService"] = service[0]
		//context.Keys["vedioService"] = service[1]
		context.Keys["vedioService"] = service[0]
		context.Next()
	}
}

//错误处理中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				context.JSON(200, gin.H{
					"code": 404,
					"msg":  fmt.Sprintf("%s", r),
				})
				context.Abort() //拦截请求：直接返回200，但响应的body中不会有数据
			}
		}()
		context.Next()
	}
}
