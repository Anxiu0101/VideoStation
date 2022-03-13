package api

import (
	"VideoStation/pkg/util"
	"VideoStation/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetVideos(c *gin.Context) {

}

func FavoriteVideo(c *gin.Context) {
	// 创建收藏视频服务
	var favoriteVideoService service.VideoService
	group := c.Query("group")

	if err := c.ShouldBind(&favoriteVideoService); err == nil {
		res := favoriteVideoService.FavoriteVideo(group)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logger().Info(err)
	}
}
