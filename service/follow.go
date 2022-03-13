package service

import (
	"VideoStation/models"
	"VideoStation/pkg/e"
	"VideoStation/pkg/util"
	"VideoStation/serializer"
)

type FollowService struct {
	UID uint `json:"uid"`
	VID uint `json:"vid"`
}

// FollowUp 关注创作者
// 1. 查询用户是否存在
// 2. 查询视频创作者是否存在
// 3. 在数据库中增添关系
// 4. 返回结果
func (service *FollowService) FollowUp() *serializer.Response {
	code := e.Success

	var user models.User
	if err := models.DB.Where("ID = ?", service.UID).Find(&user).Error; err != nil {
		util.CheckQueryErrorInDB(err)
	}

}
