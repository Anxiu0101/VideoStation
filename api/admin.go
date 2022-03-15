package api

import (
	"VideoStation/pkg/util"
	"VideoStation/service"
	"github.com/gin-gonic/gin"
)

func AdminLogin(c *gin.Context) {
	// 创建名为 用户注册服务 的 用户服务
	var userRegisterService service.UserService
	// ShouldBind 是否在此处将 JSON 反序列化了？
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.Logger().Info(err)
	}
}
