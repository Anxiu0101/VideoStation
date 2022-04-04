package serializer

import (
	"VideoStation/cache"
	"VideoStation/models"
	"time"
)

type Video struct {
	ID           uint      `json:"id" form:"id" example:"1"`
	Title        string    `json:"title" form:"title" example:"My first step in go learning"`
	Status       string    `json:"status" form:"status"`
	Video        string    `json:"video"`
	VideoType    string    `json:"video_type"`
	Introduction string    `json:"introduction"`
	CreateAt     time.Time `json:"create_at"`
	Original     bool      `json:"original"`
	Up           User      `json:"up"`
	Data         VideoData `json:"data"`
	Clicks       string    `json:"clicks"`
}

type VideoData struct {
	LikeCount     int `json:"like_count"`
	FavoriteCount int `json:"favorite_count""`
}

func BuildVideo(item models.Video, data VideoData) Video {
	clicks, _ := cache.RedisClient.Get(cache.Ctx, cache.VideoClicksKey(int(item.ID))).Result()
	return Video{
		ID:           item.ID,
		Title:        item.Title,
		Video:        item.Video,
		Introduction: item.Introduction,
		CreateAt:     item.CreatedAt,
		Up: User{
			ID:       item.User.ID,
			Username: item.User.Username,
			Avatars:  item.User.Avatars,
		},
		Data:   data,
		Clicks: clicks,
	}
}

// BuildVideosSingle 序列化视频
func BuildVideosSingle(item models.Video) Video {
	clicks, _ := cache.RedisClient.Get(cache.Ctx, cache.VideoClicksKey(int(item.ID))).Result()
	return Video{
		ID:           item.ID,
		Title:        item.Title,
		Video:        item.Video,
		Introduction: item.Introduction,
		CreateAt:     item.CreatedAt,
		Up: User{
			ID:       item.User.ID,
			Username: item.User.Username,
			Avatars:  item.User.Avatars,
		},
		Clicks: clicks,
	}

}

// BuildVideos 序列化视频列表
func BuildVideos(items []models.Video) (videos []Video) {
	for _, item := range items {
		video := BuildVideosSingle(item)
		videos = append(videos, video)
	}
	return videos
}
