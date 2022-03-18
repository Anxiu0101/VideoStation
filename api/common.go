package api

import (
	"VideoStation/conf"
	"VideoStation/pkg/e"
	"VideoStation/serializer"
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/validator.v8"
)

// ErrorResponse 返回错误信息
func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", e.Field))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return serializer.Response{
				Status: 500,
				Msg:    fmt.Sprintf("%s%s", field, tag),
				Error:  fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: e.ErrorJSONNotMatch,
			Msg:    "JSON类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}
	return serializer.Response{
		Status: e.InvalidParams,
		Msg:    "参数错误",
		Error:  fmt.Sprint(err),
	}
}
