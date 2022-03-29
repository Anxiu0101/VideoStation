package serializer

import (
	"VideoStation/models"
)

type Comment struct {
	SenderID   int `json:"sender_id"`
	ReceiverID int `json:"receiver_id"`
	VID        int `json:"vid" gorm:"column:vid"`

	Content string `json:"content" gorm:"size:255"`
}

func BuildComment(comment models.Comment) Comment {
	return Comment{
		SenderID:   comment.SenderID,
		ReceiverID: comment.ReceiverID,
		VID:        comment.VID,
		Content:    comment.Content,
	}
}
