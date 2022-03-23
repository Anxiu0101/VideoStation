package service

import (
	"VideoStation/models"
	"VideoStation/pkg/e"
	"VideoStation/pkg/logging"
	"VideoStation/pkg/util"
	"VideoStation/serializer"
)

// FavoriteVideo 收藏视频
// 1. 查询用户是否存在
// 2. 查询视频是否存在
// 3. 在数据库中增添关系
// 4. 返回结果
func (service *FavoriteVideoService) FavoriteVideo() serializer.Response {
	code := e.Success

	// 检查用户是否存在
	var user models.User
	if err := models.DB.Where("ID = ?", service.UID).Find(&user).Error; err != nil {
		return util.CheckErrorUserNoFound(err)
	}

	// 检查视频是否存在
	var video models.Video
	if err := models.DB.Where("ID = ?", service.VID).Find(&video).Error; err != nil {
		return util.CheckErrorVideoNoFound(err)
	}

	// 建立收藏对象，在数据库创建收藏关系
	data := models.Favorite{
		UID:   service.UID,
		VID:   service.VID,
		Group: service.Group,
	}
	if err := models.DB.Create(&data).Error; err != nil {
		logging.Info(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Data:   "收藏失败",
			Msg:    e.GetMsg(code),
		}
	}

	// 返回成功结果
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildFavorite(data),
		Msg:    e.GetMsg(code),
	}

}
