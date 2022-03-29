package api

import (
	"VideoStation/conf"
	"VideoStation/pkg/util"
	"VideoStation/service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func WriteComment(c *gin.Context) {
	var writeCommentService service.WriteCommentService
	if err := c.ShouldBind(&writeCommentService); err == nil {
		vid := com.StrTo(c.Param("vid")).MustInt()
		claim, _ := util.ParseToken(c.GetHeader("Authorization"))
		res := writeCommentService.Write(int(claim.ID), vid)
		c.JSON(http.StatusOK, res)
	} else {
		util.Logger().Info(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

func GetComments(c *gin.Context) {
	var showCommentsService service.ShowCommentsService
	if err := c.ShouldBind(&showCommentsService); err == nil {
		vid := com.StrTo(c.Param("vid")).MustInt()
		res := showCommentsService.Show(vid, util.GetPage(c), conf.AppSetting.PageSize)
		c.JSON(http.StatusOK, res)
	} else {
		util.Logger().Info(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
