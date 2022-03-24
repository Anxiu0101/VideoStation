package models

import (
	"gorm.io/gorm"
	"mime/multipart"
)

type Video struct {
	gorm.Model

	UID  uint `json:"uid"`
	User User `json:"user" gorm:"foreignKey:UID"`

	Title        string `json:"title" gorm:"size:50;index"`
	Video        string `json:"video" gorm:"size:255"`
	Introduction string `json:"introduction" gorm:"size:255;default:'This video has no intro'"`

	State   int     `json:"state" gorm:"default:0"`  // 视频审查，0 未审查，1 审查通过，2 审查不通过
	Clicks  int     `json:"clicks" gorm:"default:0"` // 点击量
	Weights float32 `json:"weights" gorm:"default:0"`
}

func SaveUploadFile(file, cover multipart.File, dst string) {

}

// VideoURL 返回视频地址
func (video *Video) VideoURL() string {

	return ""
}
