package service

import (
	"VideoStation/models"
	"VideoStation/pkg/e"
	"VideoStation/pkg/errorCheck"
	"VideoStation/pkg/logging"
	"VideoStation/pkg/util"
	"VideoStation/serializer"
)

// AdminService 管理员服务
type AdminService struct {
	Username string `form:"username" json:"user_name" binding:"required,min=3,max=15" example:"Anxiu"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16" example:"Anxiu123456"`
}

func (service *AdminService) Login() serializer.Response {
	code := e.Success

	// 查询管理员是否存在，是则下一步，
	//	若返回错误为 "未回应"，则返回 "用户不存在"，
	//	若返回错误不为 "未回应"，则返回 "数据库错误"
	var user models.User
	if err := models.DB.Where("username = ?", service.Username).Find(&user).Error; err != nil {
		errorCheck.CheckErrorUserNoFound(err)
	}

	// 验证管理员密码是否正确，是则下一步，否则返回 "用户密码错误"
	if !user.CheckPassword(service.Password) {
		code = e.ErrorPasswordFailCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 生成 token，是则下一步，否则返回 "Token 生成失败"
	token, err := util.GenerateToken(user.ID, service.Username, 0)
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

func (service *AdminService) BanUser() serializer.Response {
	code := e.Success

	var user models.User
	if err := models.DB.Where("username = ?", service.Username).Find(&user).Error; err != nil {
		errorCheck.CheckErrorUserNoFound(err)
	}

	// 返回成功结果
	return serializer.Response{
		Status: code,
		Data:   "The User is baned now",
		Msg:    e.GetMsg(code),
	}
}
