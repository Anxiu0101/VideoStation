package serializer

import "VideoStation/models"

type Danmu struct {
	SenderID uint `json:"sender_id"`
	VID      uint `json:"vid" gorm:"column:vid"`

	Content string `json:"content" gorm:"size:25"`
	Index   int    `json:"index"`
}

func BuildDanmu(danmu models.Danmu) Danmu {
	return Danmu{
		SenderID: danmu.SenderID,
		VID:      danmu.VID,
		Content:  danmu.Content,
		Index:    danmu.Index,
	}
}
