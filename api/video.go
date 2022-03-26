package api

import (
	"VideoStation/pkg/e"
	"VideoStation/pkg/util"
	"VideoStation/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"mime/multipart"
	"net/http"
)

// ShowVideo 获取指定视频信息
func ShowVideo(c *gin.Context) {
	var ShowVideoService service.VideoShowService
	err := c.ShouldBind(&ShowVideoService)
	if err == nil {
		vid := com.StrTo(c.Param("vid")).MustInt()
		res := ShowVideoService.Show(vid)
		c.JSON(http.StatusOK, res)
		return
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logger().Info(err)
	}
}

func Recommend(c *gin.Context) {

}

func DailyRank(c *gin.Context) {

}

// DeleteVideo 删除视频
func DeleteVideo(c *gin.Context) {
	var videoDeleteService service.VideoService
	err := c.ShouldBind(&videoDeleteService)
	if err == nil {
		vid := com.StrTo(c.Param("vid")).MustInt()
		claim, err := util.ParseToken(c.GetHeader("Authorization"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  e.ErrorUserToken,
				"error": "Token 解析失败",
			})
			return
		}
		res := videoDeleteService.DeleteVideo(uint(vid), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.Logger().Info(err)
	}
}

func Publish(c *gin.Context) {
	var videoUploadService service.VideoService
	var fileHeader = new(multipart.FileHeader)
	var err error
	fileHeader, err = c.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  e.InvalidParams,
			"error": "文件为空",
		})
		return
	}
	//c.ShouldBind(&videoUploadService)
	videoUploadService.Title = c.Query("title")
	videoUploadService.Introduction = c.Query("introduction'")
	fmt.Println(videoUploadService.Introduction + "!")
	claim, err := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  e.ErrorUserToken,
			"error": "Token 解析失败",
		})
		return
	}
	res := videoUploadService.UploadVideo(claim.ID, fileHeader, fileHeader.Size)
	c.JSON(http.StatusOK, res)
}

func UploadFile(c *gin.Context) {
	code := e.Success
	file, error := c.FormFile("file")
	if error != nil {
		code = e.ErrorUploadVideo
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    code,
			"message": e.GetMsg(code),
			"error":   code,
		})
		return
	}

	if err := c.SaveUploadedFile(file, "./upload/"+file.Filename); err != nil {
		res := ErrorResponse(err)
		c.JSON(500, res)
		return
	}

	address := "127.0.0.1:8000" + "/upload/" + file.Filename

	msg := fmt.Sprintf("%s uploaded successful", file.Filename)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"msg":     msg,
		"data":    file,
		"address": address,
	})
}
