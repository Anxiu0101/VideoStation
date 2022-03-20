package service

//func UploadFile(c *gin.Context) {
//	code := e.Success
//	file, error := c.FormFile("file")
//	if error != nil {
//		code = e.ErrorUploadVideo
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"code":    code,
//			"message": e.GetMsg(code),
//			"error":   code,
//		})
//		return
//	}
//
//	if err := c.SaveUploadedFile(file, "./upload/"+file.Filename); err != nil {
//		res := api.ErrorResponse(err)
//		c.JSON(500, res)
//		return
//	}
//
//	address := "127.0.0.1:8000" + "/upload/" + file.Filename
//
//	s := fmt.Sprintf("%s uploaded successful", file.Filename)
//	c.JSON(http.StatusOK, gin.H{
//		"code":    code,
//		"msg":     s,
//		"data":    file,
//		"address": address,
//	})
//}
