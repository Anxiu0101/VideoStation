package api

import (
	"VideoStation/pkg/util"
	"VideoStation/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FavoriteVideo(c *gin.Context) {
	// 创建收藏视频服务
	var favoriteVideoService service.FavoriteVideoService
	group := c.Query("group")

	if err := c.ShouldBind(&favoriteVideoService); err == nil {
		res := favoriteVideoService.FavoriteVideo(group)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logger().Info(err)
	}
}
