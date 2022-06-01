/**
* @Author: 18209
* @Description:
* @File:  router
* @Version: 1.0.0
* @Date: 2022/5/29 0:24
 */

package weblib

import (
	"gateway/weblib/handlers"
	"gateway/weblib/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default() //default创建的路由包含默认中间件

	//使用自定义中间件，在全局使用中间件
	ginRouter.Use(middleware.Cors(), middleware.InitMiddleware(service), middleware.ErrorMiddleware())
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))
	v1 := ginRouter.Group("/douyin/publish")
	{
		//需要登录保护
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			//投稿接口
			authed.POST("action", handlers.PublishVedio)
			//发布列表
			authed.GET("list", handlers.PublishList)
		}
	}
	ginRouter.GET("/douyin/feed", handlers.GetVedioStream)
	return ginRouter
}
