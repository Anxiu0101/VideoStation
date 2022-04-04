package models

import "gorm.io/gorm"

type History struct {
	gorm.Model

	UID  uint `json:"uid"`
	User User `json:"user" gorm:"foreignKey:UID"`

	VID   uint  `json:"vid"`
	Video Video `json:"video" gorm:"foreignKey:VID"`
}
