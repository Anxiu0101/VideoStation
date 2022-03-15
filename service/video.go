package service

import (
	"VideoStation/models"
)

type VideoService struct {
	UpID         uint         `json:"up_id"`
	Title        string       `json:"title" gorm:"size:50;index"`
	Video        models.Video `json:"video" gorm:"em"`
	Introduction string       `json:"introduction" gorm:"size:255"`
}

type FavoriteVideoService struct {
	VID  uint `json:"vid"`
	UID  uint `json:"uid"`
	UpID uint `json:"up_id"`
}

func (service *VideoService) UploadVideo(video models.Video) {

}
