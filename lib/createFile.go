package lib

import (
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

type createFileMethod interface {
	CreateFile()
}

// CreateFile - 파일 생성
func CreateFile(Name string, file *multipart.FileHeader) (*os.File, error) {
	t := time.Now()
	now := strconv.Itoa(int(t.Unix()))
	fileName := Name + now + file.Filename
	dst, err := os.Create("image/" + fileName)
	return dst, err
}
