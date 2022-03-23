package util

import (
	"io"
	"mime/multipart"
	"os"
)

func UploadToServer(file *multipart.FileHeader, fileSize int64) (int, string) {
	src, err := file.Open()
	if err != nil {
		return 500, "Can't Open file"
	}
	defer src.Close()

	out, err := os.Create("./upload/" + file.Filename)
	if err != nil {
		return 500, "Can't Create file"
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return 200, "Upload Success"
}
