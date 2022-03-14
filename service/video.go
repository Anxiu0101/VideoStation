package service

type VideoService struct {
	UpID         uint   `json:"up_id"`
	Title        string `json:"title" gorm:"size:50;index"`
	Video        string `json:"video" gorm:"size:255"`
	Introduction string `json:"introduction" gorm:"size:255"`
}

type FavoriteVideoService struct {
	VID  uint `json:"vid"`
	UID  uint `json:"uid"`
	UpID uint `json:"up_id"`
}
