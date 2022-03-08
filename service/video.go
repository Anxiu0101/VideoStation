package service

type ShowVideos struct {
	UpID         uint   `json:"publisher_id"`
	Title        string `json:"title" gorm:"size:50;index"`
	Video        string `json:"video" gorm:"size:255"`
	Introduction string `json:"introduction" gorm:"size:255"`
}
