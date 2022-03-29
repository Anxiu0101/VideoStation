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

type WriteCommentService struct {
	ReceiverID int    `json:"receiver_id"`
	Content    string `json:"content"`
}

type ShowCommentsService struct {
}

// Write 编写评论
// 1. 检查视频
// 2. 写入评论
// 3. 返回结果
func (service *WriteCommentService) Write(uid, vid int) serializer.Response {
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

	// 创建评论
	data := models.Comment{
		VID:        vid,
		SenderID:   uid,
		ReceiverID: service.ReceiverID,
		Content:    service.Content,
	}
	err := models.DB.Model(models.Comment{}).Create(&data).Error
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 返回成功结果
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildComment(data),
		Msg:    e.GetMsg(code),
	}
}

// Show 获取视频评论，
// 评论接收者为 0 则为视频评论，不为 0 则是对该数字 id 的用户的回复
// 1. 检查视频
// 2. 读取评论列表
// 3. 返回结果
func (service *ShowCommentsService) Show(vid, pageNum, pageSize int) serializer.Response {
	code := e.Success

	// 检查视频是否存在
	var video models.Video
	if err := models.DB.Where("id = ?", vid).Find(&video).Error; err != nil {
		return errorCheck.CheckErrorVideoNoFound(err)
	}

	var count int64
	if err := models.DB.Model(models.Comment{}).Where("vid = ?", vid).Count(&count).Error; err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "获取评论数失败",
		}
	}

	var comments []serializer.Comment
	err := models.DB.Model(models.Comment{}).
		Where("vid = ?", vid).
		Offset(pageNum).Limit(pageSize).
		Find(&comments).Error
	if err != nil {
		code = e.ErrorDatabase
		util.Logger().Info(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "获取评论列表失败",
		}
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 返回成功结果
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildListResponse(comments, uint(count)),
		Msg:    e.GetMsg(code),
	}
}
