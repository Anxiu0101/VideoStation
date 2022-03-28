package service

import (
	"VideoStation/models"
	"VideoStation/pkg/e"
	"VideoStation/pkg/errorCheck"
	"VideoStation/serializer"
	"fmt"
)

type CommentService struct {
	ReceiverID int    `json:"receiver_id"`
	Content    string `json:"content"`
}

// Write 编写评论
// 1. 检查视频
// 2. 写入评论
// 3. 返回结果
func (service *CommentService) Write(uid, vid int) serializer.Response {
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
	var data models.Comment
	data = models.Comment{
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

	fmt.Println("content: ", service.Content)

	// 返回成功结果
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildComment(data),
		Msg:    e.GetMsg(code),
	}
}
