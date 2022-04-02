package api

import (
	"VideoStation/conf"
	"VideoStation/pkg/util"
	"VideoStation/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminRegister(c *gin.Context) {
	var adminService service.AdminService
	if err := c.ShouldBind(&adminService); err == nil {
		res := adminService.Register()
		c.JSON(200, res)
	} else {
		util.Logger().Info(err)
		c.JSON(400, ErrorResponse(err))
	}
}

func AdminLogin(c *gin.Context) {
	var adminService service.AdminService
	// ShouldBind 是否在此处将 JSON 反序列化了？
	// ShouldBind 是绑定了URL的参数
	if err := c.ShouldBind(&adminService); err == nil {
		res := adminService.Login()
		c.JSON(200, res)
	} else {
		util.Logger().Info(err)
		c.JSON(400, ErrorResponse(err))
	}
}

func AdminBan(c *gin.Context) {
	var adminBanService service.AdminBanService
	if err := c.ShouldBind(&adminBanService); err == nil {
		res := adminBanService.BanUser()
		c.JSON(http.StatusOK, res)
	} else {
		util.Logger().Info(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// AdminVideoVerify 管理员审核视频
func AdminVideoVerify(c *gin.Context) {
	var adminVerifyService service.AdminVerifyService
	if err := c.ShouldBind(&adminVerifyService); err == nil {
		res := adminVerifyService.Verify()
		c.JSON(http.StatusOK, res)
	} else {
		util.Logger().Info(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// AdminVerifyList 管理员获取审核视频列表
func AdminVerifyList(c *gin.Context) {
	var adminVerifyListService service.AdminVerifyListService
	if err := c.ShouldBind(&adminVerifyListService); err == nil {
		res := adminVerifyListService.GetList(util.GetPage(c), conf.AppSetting.PageSize)
		c.JSON(http.StatusOK, res)
	} else {
		util.Logger().Info(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
