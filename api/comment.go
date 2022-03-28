package api

import (
	"VideoStation/pkg/util"
	"VideoStation/service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func WriteComment(c *gin.Context) {
	var writeCommentService service.CommentService
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

}
