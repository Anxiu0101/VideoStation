package models

import "gorm.io/gorm"

// Danmu 即弹幕
// 	Sender 是发送者，
// 	Content 是弹幕内容，
// 	Index 是弹幕在视频时间线的位置，单位为秒
type Danmu struct {
	gorm.Model

	SenderID uint `json:"sender_id"`
	Sender   User `json:"sender" gorm:"foreignKey:SenderID"`

	VID   uint  `json:"vid" gorm:"column:vid"`
	Video Video `json:"video" gorm:"foreignKey:VID"`

	Content string `json:"content" gorm:"size:25"`
	Index   int    `json:"index"`
}
