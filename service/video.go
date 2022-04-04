package service

import (
	"VideoStation/cache"
	"VideoStation/models"
	"VideoStation/pkg/e"
	"VideoStation/pkg/errorCheck"
	"VideoStation/pkg/logging"
	"VideoStation/pkg/util"
	"VideoStation/serializer"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type VideoService struct {
	UID  uint        `json:"uid"`
	User models.User `json:"user"`

	Title        string `json:"title" gorm:"size:50"`
	Video        string `json:"video" gorm:"size:255"`
	Introduction string `json:"introduction" gorm:"size:255"`

	State   int `json:"state"`
	Clicks  int `json:"clicks"`
	Weights int `json:"weights"`

	PageSize int `json:"page_size" form:"page_size"`
	PageNum  int `json:"page_num" form:"page_num"`
}

type VideoShowService struct {
}

type FavoriteVideoService struct {
	Group string `json:"group"`
}

// DailyRankService 每日点击排行服务
type DailyRankService struct {
}

func FavoriteAndLikeCount(vid string) (int, int) {
	var like int64
	var favorite int64
	intVid, _ := strconv.Atoi(vid)
	strLike, _ := cache.RedisClient.Get(cache.Ctx, cache.VideoLikeKey(intVid)).Result()
	strFavorite, _ := cache.RedisClient.Get(cache.Ctx, cache.VideoFavoriteKey(intVid)).Result()
	if strLike == "" || strFavorite == "" {
		//like和SQL的关键词冲突了，需要写成`like`
		models.DB.Model(&models.Interactive{}).Where("id = ? AND `like` = 1", vid).Count(&like)
		models.DB.Model(&models.Interactive{}).Where("id = ? AND favorite = 1", vid).Count(&favorite)
		//写入redis，设置6小时过期
		cache.RedisClient.Set(cache.Ctx, cache.VideoLikeKey(intVid), like, time.Hour*6)
		cache.RedisClient.Set(cache.Ctx, cache.VideoFavoriteKey(intVid), favorite, time.Hour*6)
		// count 放回类型为 int64，这里直接强转了，必有问题，但未处理
	}
	like32, _ := strconv.Atoi(strLike)
	favorite32, _ := strconv.Atoi(strFavorite)
	return like32, int(favorite32)
}

func (service *VideoShowService) Show(vid int) serializer.Response {
	code := e.Success

	var video models.Video
	err := models.DB.Model(&models.Video{}).
		Preload("User").
		Where("id = ? AND state = 1", vid).
		First(&video).Error

	if err != nil {
		code = e.InvalidParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "找不到此视频, 或视频未通过审核",
		}
	}

	like, favorite := FavoriteAndLikeCount(string(vid))

	// 收集视频数据信息
	likeStr, _ := cache.RedisClient.Get(cache.Ctx, cache.VideoLikeKey(vid)).Result()
	favoriteStr, _ := cache.RedisClient.Get(cache.Ctx, cache.VideoFavoriteKey(vid)).Result()

	var like64, favorite64 int64
	if likeStr == "" || favoriteStr == "" {
		models.DB.Model(models.Interactive{}).Where("vid = ? AND `like` = 1", vid).Count(&like64)
		models.DB.Model(models.Interactive{}).Where("vid = ? AND favorite = 1", vid).Count(&favorite64)

		cache.RedisClient.Set(cache.Ctx, cache.VideoLikeKey(vid), like64, time.Hour*6)
		cache.RedisClient.Set(cache.Ctx, cache.VideoFavoriteKey(vid), favorite64, time.Hour*6)
	}

	strClicks, _ := cache.RedisClient.Get(cache.Ctx, cache.VideoClicksKey(vid)).Result()

	if strClicks == "" {
		cache.RedisClient.RPush(cache.Ctx, cache.ClicksVideoList, vid)
		cache.RedisClient.Set(cache.Ctx, cache.VideoClicksKey(vid), video.Clicks, time.Hour*25)
	}

	cache.RedisClient.Incr(cache.Ctx, cache.VideoClicksKey(vid))

	data := serializer.VideoData{
		LikeCount:     like,
		FavoriteCount: favorite,
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildVideo(video, data),
	}
}

func (service *VideoService) DeleteVideo(vid, uid uint) serializer.Response {
	code := e.Success

	// 查询用户，是则下一步，
	// 	若返回 未回应，则返回 用户不存在
	//	若返回其他错误，则返回 数据库错误
	var user models.User
	if err := models.DB.Where("id = ?", uid).Find(&user).Error; err != nil {
		return errorCheck.CheckErrorUserNoFound(err)
	}

	// 查询视频和所属人，是则下一步，
	// 	若返回 未回应，则返回 Invalid params
	//	若返回其他错误，则返回 数据库错误
	var video models.Video
	if err := models.DB.Where("id = ? AND uid = ?", vid, uid).Find(&video).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = e.InvalidParams
			logging.Info(err)

			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		} else {
			code = e.ErrorDatabase
			logging.Info(err)

			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	}

	models.DB.Delete(&video)

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   "视频已删除",
	}
}

func (service *VideoService) UploadVideo(uid uint, file *multipart.FileHeader, fileSize int64) serializer.Response {
	code, info := models.UploadToServer(file, fileSize)
	if code != http.StatusOK {
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  info,
		}
	}

	if uid == 0 {
		code = e.InvalidParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	video := models.Video{
		UID: uid,

		Title:        service.Title,
		Video:        info,
		Introduction: service.Introduction,

		State: 0, // 未审查
	}

	// 创建视频 返回结果
	err := models.DB.Model(models.Video{}).Create(&video).Error
	if err != nil {
		code = e.ErrorDatabase
	} else {
		code = e.Success
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   "视频发布成功，待审核",
	}
}

// RankDaily 弃用
func (service *DailyRankService) RankDaily() serializer.Response {
	code := e.Success

	var videos []models.Video

	vids, err := cache.RedisClient.ZRevRange(cache.Ctx, cache.ClicksVideoList, 0, 9).Result()
	if err != nil {
		util.Logger().Info(err)
		return serializer.Response{
			Status: e.ErrorDatabase,
			Msg:    "数据库错误",
		}
	}

	print(vids)

	if len(vids) > 1 {
		order := fmt.Sprintf("FIELD(id, %s)", strings.Join(vids, ","))
		err := models.DB.Where("id in (?)", vids).Order(order).Find(&videos).Error
		if err != nil {
			return serializer.Response{
				Status: 50000,
				Msg:    "数据库连接错误",
				Error:  err.Error(),
			}
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildVideos(videos),
	}
}

func (service *VideoShowService) Rank() serializer.Response {
	code := e.Success

	var videos []models.Video
	var count int64
	models.DB.Model(&models.Video{}).Preload("User").Where("state = 1").Order("clicks desc").
		Find(&videos).Count(&count)
	for i := 0; int64(i) < count; i++ {
		tmp := strconv.Itoa(videos[i].Clicks)
		click, _ := strconv.Atoi(GetClicksFromRedis(cache.RedisClient, int(videos[i].ID), tmp))
		videos[i].Clicks = click
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildVideos(videos),
	}
}

func ClicksStoreInDB() {
	util.Logger().Info("[info]", " Clicks are stored in the database")
	var vid int          //视频id
	var key string       //redis的key
	var clicks int       //点击量数字
	var strClicks string //字符串格式
	videos := cache.RedisClient.LRange(cache.Ctx, cache.ClicksVideoList, 0, -1).Val()
	for _, i := range videos {
		vid, _ = strconv.Atoi(i)
		key = cache.VideoClicksKey(vid)
		strClicks, _ = cache.RedisClient.Get(cache.Ctx, key).Result()
		clicks, _ = strconv.Atoi(strClicks)
		//删除redis数据
		cache.RedisClient.Del(cache.Ctx, key)
		//写入数据库
		models.DB.Model(&models.Video{}).Where("id = ?", vid).Update("clicks", clicks)
	}
	//删除list
	cache.RedisClient.Del(cache.Ctx, cache.ClicksVideoList)
	util.Logger().Info("[info]", " Click volume storage completed")
}

func GetClicksFromRedis(redis *redis.Client, vid int, dbClicks string) string {
	strClicks, _ := redis.Get(cache.Ctx, cache.VideoClicksKey(vid)).Result()
	if len(strClicks) == 0 {
		//将视频ID存入点击量列表
		redis.RPush(cache.Ctx, cache.ClicksVideoList, vid)
		//将点击量存入redis并设置25小时，防止数据当天过期
		redis.Set(cache.Ctx, cache.VideoClicksKey(vid), dbClicks, time.Hour*25)
		return dbClicks
	}
	return strClicks
}
