package service

import (
	"VideoStation/models"
	"VideoStation/pkg/e"
	"VideoStation/pkg/errorCheck"
	"VideoStation/pkg/logging"
	"VideoStation/pkg/util"
	"VideoStation/serializer"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"mime/multipart"
	"net/http"
)

type VideoService struct {
	UID  uint        `json:"uid"`
	User models.User `json:"user"`

	Title        string `json:"title" gorm:"size:50"`
	Video        string `json:"video" gorm:"size:255"`
	Introduction string `json:"introduction" gorm:"size:255"`

	State   int `json:"state"`
	Clicks  int `json:"clicks"`
	Weights int `json:"weights"`

	PageSize int `json:"page_size" form:"page_size"`
	PageNum  int `json:"page_num" form:"page_num"`
}

type FavoriteVideoService struct {
	VID   uint   `json:"vid"`
	Group string `json:"group"`
}

func (service *VideoService) DeleteVideo(vid, uid uint) serializer.Response {
	code := e.Success

	// 查询用户，是则下一步，
	// 	若返回 未回应，则返回 用户不存在
	//	若返回其他错误，则返回 数据库错误
	var user models.User
	if err := models.DB.Where("id = ?", uid).Find(&user).Error; err != nil {
		return errorCheck.CheckErrorUserNoFound(err)
	}

	// 查询视频和所属人，是则下一步，
	// 	若返回 未回应，则返回 Invalid params
	//	若返回其他错误，则返回 数据库错误
	var video models.Video
	if err := models.DB.Where("id = ? AND uid = ?", vid, uid).Find(&video).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = e.InvalidParams
			logging.Info(err)

			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		} else {
			code = e.ErrorDatabase
			logging.Info(err)

			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	}

	models.DB.Delete(&video)

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   "视频已删除",
	}
}

func (service *VideoService) UploadVideo(uid uint, file *multipart.FileHeader, fileSize int64) serializer.Response {
	code, info := util.UploadToServer(file, fileSize)
	if code != http.StatusOK {
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  info,
		}
	}

	fmt.Println(service.Introduction + "!")

	video := models.Video{
		UID: uid,

		Title:        service.Title,
		Video:        info,
		Introduction: service.Introduction,

		State: 0, // 未审查
	}

	fmt.Println(video.Introduction + "!")

	// 创建视频 返回结果
	err := models.DB.Model(models.Video{}).Create(&video).Error
	if err != nil {
		code = e.ErrorDatabase
	} else {
		code = e.Success
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   "视频发布成功，待审核",
	}
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
