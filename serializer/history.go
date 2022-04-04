package serializer

import (
	"VideoStation/models"
	"time"
)

type History struct {
	UID       uint `json:"uid" form:"uid"`
	VID       uint `json:"vid" form:"vid"`
	UpdatedAt time.Time
}

func BuildHistory(history models.History) History {
	return History{
		UID:       history.UID,
		VID:       history.VID,
		UpdatedAt: history.UpdatedAt,
	}
}

// BuildHisrotys 序列化视频列表
func BuildHisrotys(items []models.History) (historys []History) {
	for _, item := range items {
		history := BuildHistory(item)
		historys = append(historys, history)
	}
	return historys
}
