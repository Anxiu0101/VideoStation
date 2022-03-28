package models

import "gorm.io/gorm"

// Comment 评论
// 	Sender 发送者，
// 	Receiver 接受者，也就是被回复者，若该字段 ID 为 0，则为视频评论
// 	Content 评论内容
type Comment struct {
	gorm.Model

	SenderID int  `json:"sender_id"`
	Sender   User `json:"sender" gorm:"foreignKey:SenderID"`

	ReceiverID int  `json:"receiver_id"`
	Receiver   User `json:"receiver" gorm:"foreignKey:ReceiverID"`

	VID   int   `json:"vid" gorm:"column:'vid'"`
	Video Video `json:"video" gorm:"foreignKey:VID"`

	Content string `json:"content" gorm:"size:255"`
}
