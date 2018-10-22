package routers

import (
	"github.com/gin-gonic/gin"
	"niurenshuo/pkg/setting"
	"niurenshuo/routers/api/v1"
)

//初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiV1 := r.Group("/api/v1")
	{
		//获取评论列表
		apiV1.GET("/comments", v1.GetComments)
		//新增评论
		apiV1.POST("/comments", v1.AddComment)
		//更新指定评论
		apiV1.PUT("/comments/:id", v1.EditComment)
		//删除指定标签
		apiV1.DELETE("/comments/:id", v1.DeleteComment)
	}

	return r
}
