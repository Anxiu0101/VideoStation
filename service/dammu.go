package service

import (
	"VideoStation/models"
	"VideoStation/pkg/e"
	"VideoStation/pkg/errorCheck"
	"VideoStation/pkg/util"
	"VideoStation/serializer"
	"errors"
	"gorm.io/gorm"
)

type SendDanmuService struct {
	Content string `json:"content" gorm:"size:25"`
	Index   int    `json:"index"` // 弹幕所在的时间点
}

type ShowDanmusService struct {
}

// Send 发送弹幕
// 1. 检查视频
// 2. 写入数据
// 3. 返回结果
func (service *SendDanmuService) Send(vid, uid int) serializer.Response {
	code := e.Success

	// 检查视频是否存在
	var video models.Video
	if err := models.DB.Where("id = ?", vid).Find(&video).Error; err != nil {
		return errorCheck.CheckErrorVideoNoFound(err)
	}

	// 检查评论内容是否为空
	if service.Content == "" {
		code = e.InvalidParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "不可以输入空的评论",
		}
	}

	data := models.Danmu{
		SenderID: uint(uid),
		VID:      uint(vid),
		Content:  service.Content,
		Index:    service.Index,
	}

	if err := models.DB.Model(models.Danmu{}).Create(&data).Error; err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "弹幕发送失败",
		}
	}

	// 返回成功结果
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildDanmu(data),
		Msg:    e.GetMsg(code),
	}
}

func (service *ShowDanmusService) Show(vid, pageNum, pageSize int) serializer.Response {
	code := e.Success

	// 检查视频是否存在
	var video models.Video
	if err := models.DB.Where("id = ?", vid).Find(&video).Error; err != nil {
		return errorCheck.CheckErrorVideoNoFound(err)
	}

	var count int64
	if err := models.DB.Model(models.Danmu{}).Where("vid = ?", vid).Count(&count).Error; err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "获取弹幕数失败",
		}
	}

	var data []serializer.Danmu
	err := models.DB.Model(models.Danmu{}).
		Where("vid = ?", vid).
		Offset(pageNum).Limit(pageSize).
		Find(&data).Error
	if err != nil {
		code = e.ErrorDatabase
		util.Logger().Info(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "获取弹幕列表失败",
		}
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "",
		}
	}

	// 返回成功结果
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildListResponse(data, uint(count)),
		Msg:    e.GetMsg(code),
	}
}
