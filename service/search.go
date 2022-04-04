package service

import (
	"VideoStation/models"
	"VideoStation/pkg/e"
	"VideoStation/pkg/util"
	"VideoStation/serializer"
	"time"
)

type SearchService struct {
	Ob string `json:"ob" form:"ob"`

	CreatedAt time.Time `json:"created_at" form:"created_at"`

	Title          string `json:"title" form:"title"`
	ClicksMaxLimit int    `json:"clicks_max_limit" form:"clicks_max_limit"`
	ClicksMinLimit int    `json:"clicks_min_limit" form:"clicks_min_limit"`

	UserName string `json:"username" form:"username"`
	Gender   uint   `json:"gender" form:"gender"`
	Age      uint   `json:"age" form:"age"`
}

func (service *SearchService) UserSearch() serializer.Response {
	code := e.Success
	db := models.DB
	users := make([]models.User, 0)

	if service.UserName != "" {
		db = db.Where("Username = ?", service.UserName)
	}

	if service.Age != 0 {
		db = db.Where("Age = ?", service.Age)
	}

	if service.Gender != 0 {
		db = db.Where("Gender = ?", service.Gender)
	}

	if err := db.Find(&users).Error; err != nil {
		code = e.ErrorDatabase
		util.Logger().Info(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   users,
	}
}

func (service *SearchService) VideoSearch() serializer.Response {
	code := e.Success
	db := models.DB
	videos := make([]models.Video, 0)

	if service.Title != "" {
		db = db.Where("Title = ?", service.Title)
	}

	if service.ClicksMinLimit != 0 {
		db = db.Where("Click > ?", service.ClicksMaxLimit)
	}

	if err := db.Find(&videos).Error; err != nil {
		code = e.ErrorDatabase
		util.Logger().Info(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   videos,
	}
}
