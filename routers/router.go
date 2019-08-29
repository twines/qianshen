package routers

import (
	"github.com/gin-gonic/gin"
	"qianshen/app/admin/v1"
	"qianshen/app/web/v1"
	"qianshen/middleware/cors"
	"qianshen/middleware/jwt"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(cors.Cors())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//后台管理
	{
		adminGroup := r.Group("/admin")
		adminApi := adminGroup.Group("/api/v1")
		adminApi.GET("/login", adminV1.Login)
		adminApi.Use(jwt.Admin())
		{
			adminApi.GET("/index", adminV1.Index)
		}
	}
	//web前端
	{
		webGroup := r.Group("/")
		webApi := webGroup.Group("/web/api/v1")
		webApi.GET("/login", webV1.Login)
		webApi.Use(jwt.Web())
		{
			webApi.GET("/index", webV1.Index)
		}
	}
	return r
}
