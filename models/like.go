package models

// Like 点赞结构体
type Like struct {
	UID uint `json:"uid" gorm:"index"`
	VID uint `json:"vid" gorm:"index"`
}
