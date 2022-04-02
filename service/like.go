package service

import (
	"VideoStation/cache"
	"VideoStation/models"
	"VideoStation/pkg/e"
	"VideoStation/pkg/errorCheck"
	"VideoStation/serializer"
	"errors"
	"gorm.io/gorm"
)

type VideoLikeService struct {
}

// LikeVideo 点赞视频
// 1. 查询视频存在
// 2. 查询用户存在
// 3. 将点赞写入缓存
// 4. 返回结果
// 这是一个使用到 redis 进行缓存，缓存的数据是总的点赞数，点赞关系通过 SQL 数据库进行查询即可。
func (service *VideoLikeService) LikeVideo(vid int, uid int) serializer.Response {
	code := e.Success

	// 检查用户是否存在
	var user models.User
	if err := models.DB.Where("id = ?", uid).Find(&user).Error; err != nil {
		return errorCheck.CheckErrorUserNoFound(err)
	}

	// 检查视频是否存在
	var video models.Video
	if err := models.DB.Where("id = ?", vid).Find(&video).Error; err != nil {
		return errorCheck.CheckErrorVideoNoFound(err)
	}

	// 从数据库中查询点赞关系
	var data models.Interactive
	err := models.DB.Model(&models.Interactive{}).Where("vid = ? AND uid = ?", vid, uid).First(&data).Error
	// 用户已为此视频点赞
	if data.Like == true {
		code = e.ErrorLikeExist
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "您已经点过赞啦",
		}
	}
	// 若找不到这个用户和该视频的关系，则创建新数据建立关系
	if errors.Is(err, gorm.ErrRecordNotFound) {
		data = models.Interactive{
			VID: vid,
			UID: uid,
		}
		models.DB.Model(&models.Interactive{}).Create(&data)
	}
	// 更新用户与视频关系
	models.DB.Model(&models.Interactive{}).Where("uid = ? AND vid = ?", uid, vid).Update("like", true)

	// 将点赞写入缓存
	strLike, _ := cache.RedisClient.Get(cache.Ctx, cache.VideoLikeKey(vid)).Result()
	if strLike != "" {
		cache.RedisClient.Incr(cache.Ctx, cache.VideoLikeKey(vid))
	}

	// 返回成功结果
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   "点赞成功",
	}
}
