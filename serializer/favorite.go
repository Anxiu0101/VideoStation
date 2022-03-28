package serializer

import "VideoStation/models"

type Favorite struct {
	UID   int    `json:"uid"`
	VID   int    `json:"vid"`
	Group string `json:"group"`
}

func BuildFavorite(interactive models.Interactive) Favorite {
	return Favorite{
		UID:   interactive.UID,
		VID:   interactive.VID,
		Group: interactive.Group,
	}
}
