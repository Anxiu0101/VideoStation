package models

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model

	UID   uint   `json:"uid"`
	VID   uint   `json:"vid"`
	Group string `json:"group" gorm:"default:'my-Favorite''"`
}
