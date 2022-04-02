package models

import (
	"VideoStation/conf"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model

	Name       string `json:"name" gorm:"size:25;default:'Admin';index;comment:管理员名"`
	Email      string `json:"email" gorm:"size:25;comment:管理员邮箱"`
	Ciphertext string `gorm:"size:255;comment:密码密文"`
	Authority  int    `json:"authority" gorm:"not null;comment:管理员权限"`
}

//SetPassword 设置密码
func (admin *Admin) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), conf.AppSetting.PassWordCost)
	if err != nil {
		return err
	}
	admin.Ciphertext = string(bytes)
	return nil
}

//CheckPassword 校验密码
func (admin *Admin) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(admin.Ciphertext), []byte(password))
	return err == nil
}
