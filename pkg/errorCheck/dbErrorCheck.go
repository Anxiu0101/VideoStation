package errorCheck

import (
	"VideoStation/pkg/e"
	"VideoStation/pkg/logging"
	"VideoStation/serializer"
	"errors"
	"gorm.io/gorm"
)

func CheckErrorUserNoFound(err error) serializer.Response {
	code := e.ErrorDatabase
	if errors.Is(err, gorm.ErrRecordNotFound) {
		code = e.ErrorNotExistUser
		logging.Info(err)

		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	} else {
		logging.Info(err)

		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
}

func CheckErrorVideoNoFound(err error) serializer.Response {
	code := e.ErrorDatabase
	if errors.Is(err, gorm.ErrRecordNotFound) {
		code = e.ErrorNotExistVideo
		logging.Info(err)

		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	} else {
		logging.Info(err)

		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
}
