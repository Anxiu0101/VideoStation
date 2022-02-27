package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	// 用户基础信息
	Username string `json:"username"`
	Password string `json:"password"`
	Avatars  string `json:"avatars"`

	// 用户详细信息
	Gender int    `json:"gender"`
	Age    int    `json:"age"`
	Email  string `json:"email"`

	//账户信息
	State int `json:"state"`
}


