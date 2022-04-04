package service

import (
	"VideoStation/models"
	"VideoStation/pkg/e"
	"VideoStation/pkg/errorCheck"
	"VideoStation/pkg/logging"
	"VideoStation/serializer"
)

// HistoryStoreInDB redis 还是不会用
func HistoryStoreInDB() {

}

type HistoryService struct {
}

type HistorySetService struct {
	VID int `json:"vid" form:"vid"`
}

func (service *HistoryService) GetHistory(uid int) serializer.Response {
	code := e.Success

	// 检查用户是否存在
	var user models.User
	if err := models.DB.Where("ID = ?", uid).Find(&user).Error; err != nil {
		return errorCheck.CheckErrorUserNoFound(err)
	}

	var historys []models.History
	if err := models.DB.Model(models.History{}).Where("UID = ?", uid).Find(&historys).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Data:   "查询历史记录失败",
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildHisrotys(historys),
	}
}

func (service *HistorySetService) SetHistory(vid, uid int) serializer.Response {
	code := e.Success

	// 检查用户是否存在
	var user models.User
	if err := models.DB.Where("ID = ?", uid).Find(&user).Error; err != nil {
		return errorCheck.CheckErrorUserNoFound(err)
	}

	// 检查视频是否存在
	var video models.Video
	if err := models.DB.Where("ID = ?", vid).Find(&video).Error; err != nil {
		return errorCheck.CheckErrorVideoNoFound(err)
	}

	if vid == 0 {
		code = e.InvalidParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	data := models.History{
		UID: uint(uid),
		VID: uint(vid),
	}

	if err := models.DB.Model(models.History{}).Create(&data).Error; err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "创建历史记录失败",
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildHistory(data),
	}
}
