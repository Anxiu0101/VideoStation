package api

import (
	"VideoStation/conf"
	"VideoStation/pkg/util"
	"VideoStation/service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func ShowDammu(c *gin.Context) {
	var showDanmusService service.ShowDanmusService
	if err := c.ShouldBind(&showDanmusService); err == nil {
		vid := com.StrTo(c.Param("vid")).MustInt()
		res := showDanmusService.Show(vid, util.GetPage(c), conf.AppSetting.PageSize)
		c.JSON(http.StatusOK, res)
	} else {
		util.Logger().Info(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func SendDammu(c *gin.Context) {
	var sendDanmuService service.SendDanmuService
	if err := c.ShouldBind(&sendDanmuService); err == nil {
		vid := com.StrTo(c.Param("vid")).MustInt()
		claim, _ := util.ParseToken(c.GetHeader("Authorization"))
		res := sendDanmuService.Send(vid, int(claim.ID))
		c.JSON(http.StatusOK, res)
	} else {
		util.Logger().Info(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
