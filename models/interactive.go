package models

import "gorm.io/gorm"

type Interactive struct {
	gorm.Model

	UID  int  `gorm:"not null"`
	User User `json:"user" gorm:"foreignKey:ID;references:UID;"`

	VID   int   `gorm:"not null;column:vid"`
	Video Video `gorm:"foreignKey:ID;references:VID;"`

	Favorite bool   `gorm:"default:false"`         //是否收藏
	Group    string `gorm:"default:'My-favorite'"` // 收藏夹
	Like     bool   `gorm:"default:false"`         //是否点赞
	//like和SQL的关键词冲突了，查询时需要写成`like`
}
