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
			loggedUserApi.GET("user/:id", api.ShowUserInfo)
		}
	}

	// 管理员操作
	adminApi := r.Group("/")
	{
		adminApi.POST("admin/login", api.AdminLogin)
	}

	// Api version-1
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		// 收藏接口
		favoriteApi := apiv1.Group("/favorite")
		{
			favoriteApi.PUT("/:uid/:vid", api.FavoriteVideo)
		}

		// 视频接口
		videoApi := apiv1.Group("/video")
		{
			videoApi.POST("/upload", api.UploadVideo)
		}
	}

	// 404 信息返回
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})

	return r
}
