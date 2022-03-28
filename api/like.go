package api

import (
	"VideoStation/pkg/util"
	"VideoStation/service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func LikeVideo(c *gin.Context) {
	var videoLikeService service.VideoLikeService
	vid := com.StrTo(c.Param("id")).MustInt()
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&videoLikeService); err == nil {
		// int 转 uint 的问题需要修复
		res := videoLikeService.LikeVideo(vid, int(claim.ID))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logger().Info(err)
	}

}
