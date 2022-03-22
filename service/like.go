package service

import (
	"VideoStation/models"
	"VideoStation/pkg/e"
	"VideoStation/pkg/util"
	"VideoStation/serializer"
)

type VideoLikeService struct {
	Username string `form:"username" json:"username"`
}

// LikeVideo 点赞视频
// 1. 查询视频存在
// 2. 查询用户存在
// 3. 将点赞写入缓存
// 4. 返回结果
func (service *VideoLikeService) LikeVideo(vid int) serializer.Response {
	code := e.Success

	// 检查用户是否存在
	var user models.User
	if err := models.DB.Where("username = ?", service.Username).Find(&user).Error; err != nil {
		return util.CheckErrorUserNoFound(err)
	}

	// 检查视频是否存在
	var video models.Video
	if err := models.DB.Where("username = ?", service.Username).Find(&video).Error; err != nil {
		return util.CheckErrorVideoNoFound(err)
	}

	// 将点赞写入缓存

	// 返回成功结果
	return serializer.Response{
		Status: code,
		Data:   "",
		Msg:    e.GetMsg(code),
	}
}
