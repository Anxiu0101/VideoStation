package routers

import (
	"VideoStation/api"
	"VideoStation/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "pong!",
		})
	})

	userApi := r.Group("/")
	{
		// 用户登录与用户注册
		unLogUserApi := userApi.Group("/user")
		{
			unLogUserApi.POST("/register", api.UserRegister)
			unLogUserApi.POST("/login", api.UserLogin)
		}

		// 用户其他操作
		loggedUserApi := userApi.Group("/")
		loggedUserApi.Use(middleware.JWT())
		{
			loggedUserApi.GET("user/:id", api.ShowUserInfo)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})

	return r
}
