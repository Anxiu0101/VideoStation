package controller

import "github.com/gin-gonic/gin"

// GetUserInfo
// @Summary 获取用户详细信息
// @Produce  json
// @Param token query string true "token"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /user/space [get]
func GetUserInfo(c *gin.Context) {
	//username = c.Query("username")
	//password = c.Query("password")

}
