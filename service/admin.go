package service

import (
	"VideoStation/models"
	"VideoStation/pkg/e"
	"VideoStation/pkg/errorCheck"
	"VideoStation/pkg/logging"
	"VideoStation/pkg/util"
	"VideoStation/serializer"
	"errors"
	"gorm.io/gorm"
)

// AdminService 管理员服务
type AdminService struct {
	Username string `form:"username" json:"user_name"`
	Password string `form:"password" json:"password"`
}

type AdminBanService struct {
	UID int `json:"uid" form:"uid"`
}

type AdminVerifyService struct {
	State int `json:"state" form:"state"`
	VID   int `json:"vid" form:"vid"`
}

type AdminVerifyListService struct {
}

func (service *AdminService) Register() serializer.Response {
	code := e.Success

	var admin models.Admin
	count := models.DB.Model(models.Admin{}).Where("Name = ?", service.Username).Find(&admin).RowsAffected
	if count > 0 {
		code = e.ErrorHaveExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	admin.Name = service.Username
	if err := admin.SetPassword(service.Password); err != nil {
		code = e.ErrorFailEncryption
		util.Logger().Info(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if err := models.DB.Create(&admin).Error; err != nil {
		code = e.ErrorDatabase
		util.Logger().Info(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *AdminService) Login() serializer.Response {
	code := e.Success

	// 查询管理员是否存在，是则下一步，
	//	若返回错误为 "未回应"，则返回 "用户不存在"，
	//	若返回错误不为 "未回应"，则返回 "数据库错误"
	var admin models.Admin
	if err := models.DB.Where("name = ?", service.Username).Find(&admin).Error; err != nil {
		errorCheck.CheckErrorUserNoFound(err)
	}

	// 验证管理员密码是否正确，是则下一步，否则返回 "用户密码错误"
	if !admin.CheckPassword(service.Password) {
		code = e.ErrorPasswordFailCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 生成 token，是则下一步，否则返回 "Token 生成失败"
	token, err := util.GenerateToken(admin.ID, service.Username, 0)
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
		Data:   serializer.TokenData{User: serializer.BuildAdmin(admin), Token: token},
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

// Verify 视频审核
// 1. 查询视频存在且为被审核
// 2. 更新视频状态
// 3. 返回成功结果
func (service *AdminVerifyService) Verify() serializer.Response {
	code := e.Success

	if service.State == 0 || service.VID == 0 {
		code = e.InvalidParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "传入了空值，参数错误",
		}
	}

	// 没做处理，当视频存在但是已被其他管理员审核通过时，返回状态是视频不存在
	var video models.Video
	if err := models.DB.Where("id = ? AND State = 0", service.VID).Find(&video).Error; err != nil {
		return errorCheck.CheckErrorVideoNoFound(err)
	}

	// 更新字段
	if err := models.DB.Model(&video).Where("id = ?", service.VID).Update("State", service.State).Error; err != nil {
		code := e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "视频状态更新失败",
		}
	}

	// 返回成功结果
	return serializer.Response{
		Status: code,
		Data:   "视频审核通过成功",
		Msg:    e.GetMsg(code),
	}
}

// GetList 获取待审核视频列表
// 1. 获取总数
// 2. 获取所有字段 State 值为 0 的视频
// 3. 返回成功结果
func (service *AdminVerifyListService) GetList(pageNum, pageSize int) serializer.Response {
	code := e.Success

	var count int64
	if err := models.DB.Model(models.Video{}).Where("State = 0").Count(&count).Error; err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "获取待审核视频数失败",
		}
	}

	var videos []models.Video
	if err := models.DB.Model(models.Video{}).Preload("User").Offset(pageNum).Limit(pageSize).Find(&videos).Error; err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "获取待审核视频列表失败",
		}
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "No video is waiting verify",
		}
	}

	// 返回成功结果
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildListResponse(videos, uint(count)),
	}
}
