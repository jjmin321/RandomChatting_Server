package controller

import (
	"RandomChatting_Server/model"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

// PutImage - 프로필 사진 등록
func PutImage(c echo.Context) error {
	Name := c.Get("Name").(string)
	Pw := c.Get("Pw").(string)
	file, err := c.FormFile("image")
	t := time.Now()
	now := strconv.Itoa(int(t.Unix()))
	fileName := Name + now + file.Filename
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"status":  400,
			"message": "파일을 읽는 데 실패하였습니다. 다시 시도해주세요.",
		})
	}
	src, err := file.Open()
	defer src.Close()
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "파일을 여는 데 실패하였습니다. 다시 시도해주세요.",
		})
	}
	dst, err := os.Create("image/" + fileName)
	defer dst.Close()
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "파일을 생성하는 데 실패하였습니다. 다시 시도해주세요.",
		})
	}
	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "파일을 저장하는 데 실패하였습니다. 다시 시도해주세요.",
		})
	}
	err = model.UpdateImage(Name, Pw, fileName)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "데이터베이스에 저장하는 데 실패하였습니다. 다시 시도해주세요.",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "프로필 사진 등록에 성공하셨습니다.",
	})
}
