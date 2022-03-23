package serializer

import (
	"VideoStation/models"
)

type User struct {
	ID       uint   `json:"id" form:"id" example:"1"`
	Username string `json:"user_name" form:"user_name" example:"FanOne"`
	Avatars  string `json:"avatars"`

	Status   string `json:"status" form:"status"`
	CreateAt int64  `json:"create_at" form:"create_at"`
}

//BuildUser 序列化用户
func BuildUser(user models.User) User {
	return User{
		ID:       user.ID,
		Username: user.Username,
		CreateAt: user.CreatedAt.Unix(),
	}
}

type UserInfo struct {
	Username string `form:"username" json:"username"`
	Avatars  string `form:"avatars" json:"avatars"`
	Gender   uint   `form:"gender" json:"gender"`
	Age      uint   `form:"age" json:"age"`
	Email    string `form:"email" json:"email"`
}

func BuildUserInfo(user models.User) UserInfo {
	return UserInfo{
		Username: user.Username,
		Avatars:  user.Avatars,
		Gender:   user.Gender,
		Age:      user.Age,
		Email:    user.Email,
	}
}
