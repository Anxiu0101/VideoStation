package api

import (
	"VideoStation/pkg/util"
	"VideoStation/serializer"
	"VideoStation/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Search(c *gin.Context) {
	var service service.SearchService
	if err := c.ShouldBind(&service); err == nil {
		obj := service.Ob
		if obj == "User" || obj == "user" {
			res := service.UserSearch()
			c.JSON(http.StatusOK, res)
		} else if obj == "Video" || obj == "video" {
			res := service.VideoSearch()
			c.JSON(http.StatusOK, res)
		} else {
			res := serializer.Response{
				Status: http.StatusBadRequest,
				Data:   "查询对象错误",
			}
			c.JSON(http.StatusBadRequest, res)
		}
	} else {
		util.Logger().Info(err)
		c.JSON(400, ErrorResponse(err))
	}
}
