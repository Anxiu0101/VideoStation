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
		loggedUserApi := userApi.Group("/user")
		loggedUserApi.Use(middleware.JWT())
		{
			loggedUserApi.GET("/:id/info", api.ShowUserInfo)
			loggedUserApi.POST("/:id/info", api.UpdateUserInfo)
			loggedUserApi.POST("/:id/avatars", api.UploadFile)
			loggedUserApi.POST("/:id/password", api.ResetPassword)
		}
	}

	// 管理员操作
	adminApi := r.Group("/")
	{
		adminApi.POST("/admin/login", api.AdminLogin)
		adminApi.POST("/admin/register")
		adminApi.POST("/admin/ban", api.AdminBan)
		adminApi.POST("/admin/verify", api.AdminVideoVerify)
		adminApi.GET("/admin/verify/list", api.AdminVerifyList)
	}

	// Api version one
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		// 视频接口
		videoApi := apiv1.Group("/")
		{
			videoApi.GET("/videos", api.Recommend)
			videoApi.GET("/videos/rank", api.DailyRank)
			videoApi.GET("/video/:vid", api.ShowVideo)
			videoApi.DELETE("/video/:vid", api.DeleteVideo)
			videoApi.POST("/video/upload", api.Publish)

			videoApi.POST("/video/:vid/favorite", api.FavoriteVideo) // 用户收藏
			videoApi.POST("/video/:vid/like", api.LikeVideo)         // 用户点赞

			// 评论功能
			// 设定了一个默认用户，id为 0，当用户评论视频时，receiver id 就是 0。若用户回复某位用户的评论时，receiver id 就是被回复的用户
			videoApi.GET("/video/:vid/comments", api.GetComments)  // 查看评论
			videoApi.POST("/video/:vid/comment", api.WriteComment) // 用户评论

			// 弹幕功能
			videoApi.GET("/video/:vid/dammu", api.ShowDammu)  // 获取视频弹幕
			videoApi.POST("/video/:vid/dammu", api.SendDammu) // 发送弹幕
		}
	}

	// 404 信息返回
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})

	return r
}
