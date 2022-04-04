package api

import (
	"VideoStation/pkg/util"
	"VideoStation/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHistory(c *gin.Context) {
	var service service.HistoryService
	if err := c.ShouldBind(&service); err == nil {
		claim, _ := util.ParseToken(c.GetHeader("Authorization"))
		res := service.GetHistory(int(claim.ID))
		c.JSON(http.StatusOK, res)
	} else {
		util.Logger().Info(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func SetHistory(c *gin.Context) {
	var service service.HistorySetService
	if err := c.ShouldBind(&service); err == nil {
		claim, _ := util.ParseToken(c.GetHeader("Authorization"))
		res := service.SetHistory(service.VID, int(claim.ID))
		c.JSON(http.StatusOK, res)
	} else {
		util.Logger().Info(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
