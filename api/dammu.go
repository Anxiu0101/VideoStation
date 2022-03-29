package api

import (
	"VideoStation/pkg/util"
	"VideoStation/service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func ShowDammu(c *gin.Context) {
	var showDammusService service.ShowDammusService
	if err := c.ShouldBind(&showDammusService); err == nil {
		vid := com.StrTo(c.Param("vid")).MustInt()
		res := showDammusService.Show(vid)
		c.JSON(http.StatusOK, res)
	} else {
		util.Logger().Info(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func SendDammu(c *gin.Context) {
	var sendDammuService service.SendDammuService
	if err := c.ShouldBind(&sendDammuService); err == nil {
		vid := com.StrTo(c.Param("vid")).MustInt()
		claim, _ := util.ParseToken(c.GetHeader("Authorization"))
		res := sendDammuService.Send(vid, int(claim.ID))
		c.JSON(http.StatusOK, res)
	} else {
		util.Logger().Info(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
