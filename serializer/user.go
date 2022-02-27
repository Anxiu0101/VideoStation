package serializer

import "VideoStation/models"

type User struct {
	ID       uint   `json:"id" form:"id" example:"1"`
	Username string `json:"user_name" form:"user_name" example:"FanOne"`
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
