package api

import (
	"VideoStation/pkg/e"
	"VideoStation/pkg/util"
	"VideoStation/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetVideo(c *gin.Context) {

}

func Recommend(c *gin.Context) {

}

func Publish(c *gin.Context) {
	var videoUploadService service.VideoService
	_, fileHeader, _ := c.Request.FormFile("file")
	c.ShouldBind(&videoUploadService)
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
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

	s := fmt.Sprintf("%s uploaded successful", file.Filename)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"msg":     s,
		"data":    file,
		"address": address,
	})
}
