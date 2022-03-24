package api

import (
	"VideoStation/pkg/util"
	"VideoStation/service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func FavoriteVideo(c *gin.Context) {
	// 创建收藏视频服务
	var favoriteVideoService service.FavoriteVideoService

	if err := c.ShouldBind(&favoriteVideoService); err == nil {
		favoriteVideoService.VID = uint(com.StrTo(c.Param("id")).MustInt())
		claim, _ := util.ParseToken(c.GetHeader("Authorization"))
		res := favoriteVideoService.FavoriteVideo(claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logger().Info(err)
	}
}
