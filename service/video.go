package service

import (
	"VideoStation/models"
)

type VideoService struct {
	UpID         uint   `json:"up_id"`
	Title        string `json:"title" gorm:"size:50;index"`
	Video        string `json:"video" gorm:"size:255"`
	Introduction string `json:"introduction" gorm:"size:255"`
}

type FavoriteVideoService struct {
	VID   uint   `json:"vid"`
	UID   uint   `json:"uid"`
	Group string `json:"group"`
}

func (service *VideoService) UploadVideo(video models.Video) {

}

//func FilesController(c *gin.Context) {
//	file, err := c.FormFile("raw")
//	if err != nil {
//		log.Fatal(err)
//	}
//	exe, err := os.Executable()
//	if err != nil {
//		log.Fatal(err)
//	}
//	dir := filepath.Dir(exe)
//	if err != nil {
//		log.Fatal(err)
//	}
//	filename := uuid.New().String()
//	uploads := filepath.Join(dir, "uploads")
//	err = os.MkdirAll(uploads, os.ModePerm)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fullpath := path.Join("uploads", filename+filepath.Ext(file.Filename))
//	fileErr := c.SaveUploadedFile(file, filepath.Join(dir, fullpath))
//	if fileErr != nil {
//		log.Fatal(fileErr)
//	}
//	c.JSON(http.StatusOK, gin.H{"url": "/" + fullpath})
//}
