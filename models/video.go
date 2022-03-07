package models

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model

	UpID uint `json:"publisher_id"`
	Up   User `json:"publisher" gorm:"foreignKey:UpID"`

	Title        string `json:"title" gorm:"size:50;index"`
	Video        string `json:"video" gorm:"size:255"`
	Introduction string `json:"introduction" gorm:"size:255"`

	// 视频审查，0 未审查，1 审查通过，2 审查不通过
	State int `json:"state" gorm:"default:0"`
	// 点击量
	Clicks int `json:"clicks" gorm:"default:0"`
}
