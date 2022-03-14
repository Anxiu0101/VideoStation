package api

import (
	"VideoStation/service"
	"github.com/gin-gonic/gin"
)

func GetVideos(c *gin.Context) {

}

func UploadVideo(c *gin.Context) {
	// 创建视频上传服务
	var UploadVideoService service.VideoService

}
