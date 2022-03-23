package service

import (
	"VideoStation/cache"
	"VideoStation/models"
	"VideoStation/pkg/e"
	"VideoStation/pkg/util"
	"VideoStation/serializer"
)

type VideoLikeService struct {
}

// LikeVideo 点赞视频
// 1. 查询视频存在
// 2. 查询用户存在
// 3. 将点赞写入缓存
// 4. 返回结果
func (service *VideoLikeService) LikeVideo(vid int, uid uint) serializer.Response {
	code := e.Success

	// 检查用户是否存在
	var user models.User
	if err := models.DB.Where("id = ?", uid).Find(&user).Error; err != nil {
		return util.CheckErrorUserNoFound(err)
	}

	// 检查视频是否存在
	var video models.Video
	if err := models.DB.Where("id = ?", vid).Find(&video).Error; err != nil {
		return util.CheckErrorVideoNoFound(err)
	}

	// 将点赞写入缓存
	var data models.Interactive
	models.DB.Model(&models.Interactive{}).Where("vid = ? AND uid = ?", vid, uid).First(&data)
	if data.Like == true {
		code = e.ErrorLikeExist
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	models.DB.Model(&models.Interactive{}).Where("uid = ? AND vid = ?", uid, vid).Update("like", true)
	strLike, _ := cache.RedisClient.Get(cache.Ctx, cache.VideoLikeKey(vid)).Result()
	if strLike != "" {
		cache.RedisClient.Incr(cache.Ctx, cache.VideoLikeKey(vid))
	}

	// 返回成功结果
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
