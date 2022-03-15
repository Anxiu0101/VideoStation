package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model

	Name      string `json:"name" gorm:"size:25;default:'Admin';comment:管理员名"`
	Email     string `json:"email" gorm:"size:25;comment:管理员邮箱"`
	Password  string `json:"password" gorm:"size:255;not null;comment:密码"`
	Authority int    `json:"authority" gorm:"not null;comment:管理员权限"`
}
