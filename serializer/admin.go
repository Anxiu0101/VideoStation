package serializer

import "VideoStation/models"

func BuildAdmin(admin models.Admin) User {
	return User{
		ID:       admin.ID,
		Username: admin.Name,
		CreateAt: admin.CreatedAt.Unix(),
	}
}
