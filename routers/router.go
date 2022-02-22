package routers

import (
	"VideoStation/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	userApi := r.Group("/")
	{
		userApi.POST("/user/space", controller.GetUserInfo)
	}

	return r
}
