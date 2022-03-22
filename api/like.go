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
	if err := c.ShouldBind(&videoLikeService); err == nil {
		res := videoLikeService.LikeVideo(vid)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logger().Info(err)
	}

}
