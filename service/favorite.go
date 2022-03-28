package service

import (
	"VideoStation/cache"
	"VideoStation/models"
	"VideoStation/pkg/e"
	"VideoStation/pkg/errorCheck"
	"VideoStation/pkg/logging"
	"VideoStation/serializer"
	"errors"
	"gorm.io/gorm"
)

// FavoriteVideo 收藏视频
// 1. 查询用户是否存在
// 2. 查询视频是否存在
// 3. 在数据库中增添关系
// 4. 返回结果
func (service *FavoriteVideoService) FavoriteVideo(uid, vid int) serializer.Response {
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

	// 检查收藏关系是否已存在
	var data models.Interactive
	err := models.DB.Model(&models.Interactive{}).Where("v_id = ? AND uid = ?", vid, uid).First(&data).Error
	// 用户已收藏该视频，且收藏组未变更
	if data.Favorite == true && data.Group == service.Group {
		code = e.ErrorFavoriteExist
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "该视频已被收藏",
		}
	}
	// 若该用户与视频未建立关系，则建立
	if errors.Is(err, gorm.ErrRecordNotFound) {
		data = models.Interactive{
			VID: vid,
			UID: uid,
		}
		if err := models.DB.Model(&models.Interactive{}).Create(&data); err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Data:   "收藏失败",
				Msg:    e.GetMsg(code),
			}
		}
	}

	if service.Group == "" {
		service.Group = "My-favorite"
	}
	data = models.Interactive{
		VID:      vid,
		UID:      uid,
		Favorite: true,
		Group:    service.Group,
	}

	// 更新用户与视频关系
	models.DB.Model(models.Interactive{}).Where("uid = ? AND v_id = ?", uid, vid).Updates(&data)

	strFavorite, _ := cache.RedisClient.Get(cache.Ctx, cache.VideoFavoriteKey(vid)).Result()
	print("strFavorite: " + strFavorite)
	if strFavorite != "" {
		cache.RedisClient.Incr(cache.Ctx, cache.VideoFavoriteKey(vid))
	}

	// 返回成功结果
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildFavorite(data),
		Msg:    e.GetMsg(code),
	}

}
