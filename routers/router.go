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

	// 用户操作
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
			loggedUserApi.GET("/user/:id/info", api.ShowUserInfo)
			loggedUserApi.POST("/user/:id/info", api.UpdateUserInfo)
			loggedUserApi.POST("/user/:id/avatars", api.UploadFile)
			loggedUserApi.POST("/user/:id/password", api.ResetPassword)
		}
	}

	// 管理员操作
	adminApi := r.Group("/")
	{
		adminApi.POST("/admin/login", api.AdminLogin)
	}

	// Api version one
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		// 视频接口
		videoApi := apiv1.Group("/")
		{
			videoApi.GET("/video/:vid", api.GetVideo)
			videoApi.POST("/video/upload", api.Publish)
			videoApi.GET("/video/:vid/favorite", api.FavoriteVideo)
			videoApi.POST("/video/:vid/like", api.LikeVideo)
		}
	}

	// 404 信息返回
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})

	return r
}
