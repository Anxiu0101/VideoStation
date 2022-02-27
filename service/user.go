package service

import (
	"VideoStation/models"
	"VideoStation/pkg/e"
	_ "VideoStation/pkg/logging"
	"VideoStation/serializer"
)

// UserService 用户服务
type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" example:"FanOne"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16" example:"FanOne666"`
}

func (service *UserService) Register() *serializer.Response {
	code := e.Success
	var user models.User
	var count int64
	models.DB.Model(&models.User{}).Where("user_name = ?", service.UserName).Find(&user).Count(&count)

	/* Validation */
	if count == 1 {
		code = e.ErrorHaveExistUser
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	//表单验证
	//if count == 1 {
	//	code = e.ErrorExistUser
	//	return &serializer.Response{
	//		Status: code,
	//		Msg:    e.GetMsg(code),
	//	}
	//}
	//user.UserName = service.UserName
	////加密密码
	//if err := user.SetPassword(service.Password); err != nil {
	//	logging.Info(err)
	//	code = e.ErrorFailEncryption
	//	return &serializer.Response{
	//		Status: code,
	//		Msg:    e.GetMsg(code),
	//	}
	//}
	////创建用户
	//if err := model.DB.Create(&user).Error; err != nil {
	//	logging.Info(err)
	//	code = e.ErrorDatabase
	//	return &serializer.Response{
	//		Status: code,
	//		Msg:    e.GetMsg(code),
	//	}
	//}
	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
