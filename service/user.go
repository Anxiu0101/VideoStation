package service

import (
	"VideoStation/models"
	"VideoStation/pkg/e"
	"VideoStation/pkg/logging"
	_ "VideoStation/pkg/logging"
	"VideoStation/serializer"
)

// UserService 用户服务
type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" example:"FanOne"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16" example:"FanOne666"`
}

// Register 用户注册，分为三个步骤，
// 1. 先查询用户名是否已存在，是则下一步，否则返回 "用户已存在"
// 2. 为新用户设置密码，成功则下一步，否则返回 "加密失败"
// 3. 在数据库中创建新用户，成功则返回成功，否则返回 "数据库错误"
func (service *UserService) Register() *serializer.Response {
	code := e.Success
	var user models.User
	var count int64
	models.DB.Model(&models.User{}).Where("user_name = ?", service.UserName).Find(&user).Count(&count)

	// 查询用户名是否已存在，是则下一步，否则返回 "用户已存在"
	if count > 0 {
		code = e.ErrorHaveExistUser
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 为新用户设置密码，成功则下一步，否则返回 "加密失败"
	user.Username = service.UserName
	if err := user.SetPassword(service.Password); err != nil {
		logging.Info(err)
		code = e.ErrorFailEncryption
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 在数据库中创建新用户，成功则返回成功，否则返回 "数据库错误"
	if err := models.DB.Create(&user).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
