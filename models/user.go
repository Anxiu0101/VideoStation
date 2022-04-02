package models

import (
	"VideoStation/conf"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	// 用户基础信息
	Username   string `json:"username" gorm:"uniqueIndex;size:25;not null"`
	Ciphertext string `gorm:"size:255"`
	Avatars    string `json:"avatars"`

	// 用户详细信息
	Gender uint   `json:"gender"`
	Age    uint   `json:"age"`
	Email  string `json:"email" gorm:"size:25;not null;index"`

	//账户信息
	State int `json:"state" gorm:"default:0;size:5;comment:'0为正常.1为封禁'"`
}

//SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), conf.AppSetting.PassWordCost)
	if err != nil {
		return err
	}
	user.Ciphertext = string(bytes)
	return nil
}

//CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Ciphertext), []byte(password))
	return err == nil
}
