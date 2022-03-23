package models

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model

	UID uint `json:"uid"`

	VID   uint  `json:"vid"`
	Video Video `json:"video" gorm:"foreignKey:VID"`

	Group string `json:"group" gorm:"default:'My-Favorite'"`
}

// 还未重构，重构后将会将 favorite 集成到 Interactive 结构体中，重构后收藏和点赞功能不再设置单独的结构体
