package models

import "gorm.io/gorm"

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
