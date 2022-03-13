package serializer

import "VideoStation/models"

type Favorite struct {
	ID    uint   `json:"id"`
	UID   uint   `json:"uid"`
	VID   uint   `json:"vid"`
	Group string `json:"group"`
}

func BuildFavorite(favorite models.Favorite) Favorite {
	return Favorite{
		UID:   favorite.UID,
		VID:   favorite.VID,
		Group: favorite.Group,
	}
}
