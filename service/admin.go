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

type AdminBanService struct {
	UID int `json:"uid" form:"uid"`
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

// BanUser 封禁用户
// 1. 检查用户存在
// 2. 设置用户状态
// 3. 返回成功结果
func (service *AdminBanService) BanUser() serializer.Response {
	code := e.Success

	print(service.UID)
	if service.UID == 0 {
		code = e.InvalidParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "不可以封禁默认用户",
		}
	}

	var user models.User
	if err := models.DB.Where("id = ?", service.UID).Find(&user).Error; err != nil {
		return errorCheck.CheckErrorUserNoFound(err)
	}

	if err := models.DB.Model(models.User{}).Where("id = ?", service.UID).Update("State", 1).Error; err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "封禁失败",
		}
	}

	// 返回成功结果
	return serializer.Response{
		Status: code,
		Data:   "The User is baned now",
		Msg:    e.GetMsg(code),
	}
}
