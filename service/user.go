package service

import (
	"errors"

	"gorm.io/gorm"

	"VideoStation/models"
	"VideoStation/pkg/e"
	"VideoStation/pkg/logging"
	"VideoStation/pkg/util"
	"VideoStation/serializer"
)

// UserService 用户服务
type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" example:"FanOne"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16" example:"FanOne666"`
}

// Register 用户注册，
// 1. 先查询用户名是否已存在
// 2. 为新用户设置密码
// 3. 在数据库中创建新用户
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

// Login 用户登录，
// 1. 查询用户
// 2. 验证用户密码
// 3. 生成 token
// 4. 返回结果
func (service *UserService) Login() serializer.Response {
	code := e.Success
	var user models.User

	// 查询用户是否存在，是则下一步，
	//	若返回错误为 "未回应"，则返回 "用户不存在"，
	//	若返回错误不为 "未回应"，则返回 "数据库错误"
	if err := models.DB.Where("user_name=?", service.UserName).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logging.Info(err)
			code = e.ErrorNotExistUser

			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		if err != nil {
			logging.Info(err)
			code = e.ErrorDatabase

			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	}

	// 验证用户密码是否正确，是则下一步，否则返回 "用户密码错误"
	if user.CheckPassword(service.Password) == false {
		code = e.ErrorPasswordFailCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 生成 token，是则下一步，否则返回 "Token 生成失败"
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		logging.Info(err)
		code = e.ErrorUserToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 返回成功结果
	return serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    e.GetMsg(code),
	}
}