package controller

import (
	"RandomChatting_Server/lib"
	"RandomChatting_Server/model"
	"io"

	"github.com/labstack/echo"
)

type putImageMethod interface {
	PutImage()
}

// PutImage - 프로필 사진 등록
func PutImage(c echo.Context) error {
	Name := c.Get("Name").(string)
	Pw := c.Get("Pw").(string)
	file, err := c.FormFile("image")
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
	dst, fileName, err := lib.CreateFile(Name, file)
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
