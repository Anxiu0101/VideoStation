package routers

import (
	"VideoStation/controller"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userApi := r.Group("/")
	{
		userApi.POST("/user/space", controller.GetUserInfo)
	}

	return r
}
