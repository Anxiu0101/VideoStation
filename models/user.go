package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	// 用户基础信息
	Username   string `json:"username"`
	Ciphertext string // 加密密码
	Avatars    string `json:"avatars"`

	// 用户详细信息
	Gender int    `json:"gender"`
	Age    int    `json:"age"`
	Email  string `json:"email"`

	//账户信息
	State int `json:"state"`
}

const (
	PassWordCost = 12 //密码加密难度
)

//SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
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
