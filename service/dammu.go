package service

import (
	"VideoStation/pkg/e"
	"VideoStation/serializer"
)

type SendDammuService struct {
	Content string `json:"content" gorm:"size:25"`
	Index   int    `json:"index"` // 弹幕所在的时间点
}

type ShowDammusService struct {
}

func (service *SendDammuService) Send(vid, uid int) serializer.Response {
	code := e.Success

	// 返回成功结果
	return serializer.Response{
		Status: code,
		Data:   "",
		Msg:    e.GetMsg(code),
	}
}

func (service *ShowDammusService) Show(vid int) serializer.Response {
	code := e.Success

	// 返回成功结果
	return serializer.Response{
		Status: code,
		Data:   "",
		Msg:    e.GetMsg(code),
	}
}
