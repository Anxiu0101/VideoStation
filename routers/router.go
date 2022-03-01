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
		// 用户登录与用户注册
		unLogUserApi := userApi.Group("/user")
		{
			unLogUserApi.POST("/register", controller.UserRegister)
			unLogUserApi.POST("/login", controller.UserLogin)
		}

		// 用户其他操作
		userApi.GET("user/:id")

	}

	return r
}
