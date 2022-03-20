package api

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"

	"VideoStation/pkg/util"
	"VideoStation/service"
)

// UserRegister
// @Tags USER
// @Summary 用户注册
// @Produce json
// @Accept json
// @Param data body service.UserService true "用户名, 密码"
// @Success 200 {object} serializer.ResponseUser "{"status":200,"data":{},"msg":"ok"}"
// @Failure 500  {object} serializer.ResponseUser "{"status":500,"data":{},"Msg":{},"Error":"error"}"
// @Router /user/register [post]
func UserRegister(c *gin.Context) {
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

// UserLogin
// @Tags USER
// @Summary 用户登录
// @Produce json
// @Accept json
// @Param     data    body     service.UserService    true      "user_name, password"
// @Success 200 {object} serializer.ResponseUser "{"success":true,"data":{},"msg":"登陆成功"}"
// @Failure 500 {object} serializer.ResponseUser "{"status":500,"data":{},"Msg":{},"Error":"error"}"
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	// 创建名为 用户登录服务 的 用户服务
	var userLoginService service.UserService
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.Login()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.Logger().Info(err)
	}
}

// ShowUserInfo
// @Tags USER
// @Summary 用户资料展示
// @Produce json
// @Accept json
// @Param     data    body     service.UserService    true      "user_name, password"
// @Success 200 {object} serializer.ResponseUser "{"success":true,"data":{},"msg":"ok"}"
// @Failure 500 {object} serializer.ResponseUser "{"status":500,"data":{},"Msg":{},"Error":"error"}"
// @Router /user/:id [get]
func ShowUserInfo(c *gin.Context) {
	// 创建名为 用户信息获取服务
	var userShowInfoService service.UserService
	if err := c.ShouldBind(&userShowInfoService); err == nil {
		id := com.StrTo(c.Param("id")).MustInt()
		res := userShowInfoService.GetUserInfo(id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logger().Info(err)
	}
}

//func ResetPassword(c *gin.Context) {
//	var userResetPasswordService service.UserService
//	if err := c.ShouldBind(&userResetPasswordService); err != nil {
//		newPassword := c.
//	}
//}
//
//func ()  {
//
//}
