package models

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model

	UpID uint `json:"publisher_id"`
	Up   User `json:"publisher" gorm:"foreignKey:UpID"`

	Video       string    `json:"video"`
	DanmuList   []Danmu   `json:"danmu_list"`
	CommentList []Comment `json:"comment_list"`
	State       int       `json:"state"`
}

// Danmu 即弹幕
// 	Sender 是发送者，
// 	Content 是弹幕内容，
// 	Index 是弹幕在视频时间线的位置，单位为秒
type Danmu struct {
	gorm.Model

	SenderID uint `json:"sender_id"`
	Sender   User `json:"sender" gorm:"foreignKey:SenderID"`

	Content string `json:"content"`
	Index   int    `json:"index"`
}

// Comment 评论
// 	Sender 发送者，
// 	Receiver 接受者，也就是被回复者，若该字段 ID 为 0，则为视频评论
// 	Content 评论内容
type Comment struct {
	gorm.Model

	SenderID uint `json:"sender_id"`
	Sender   User `json:"sender" gorm:"foreignKey:SenderID"`

	ReceiverID uint `json:"receiver_id"`
	Receiver   User `json:"receiver" gorm:"foreignKey:ReceiverID"`

	Content string `json:"content"`
}
