package controller

import (
	"RandomChatting_Server/model"

	"github.com/labstack/echo"
)

// PatchNameParam - 파라미터 형식 정의 구조체
type PatchNameParam struct {
	Name string `json:"name" form:"name" query:"name"`
}

// PatchName - 내 정보 변경 메서드
func PatchName(c echo.Context) error {
	ID := c.Get("ID").(string)
	u := new(PatchNameParam)
	if err := c.Bind(u); err != nil {
		return err
	}
	err := model.UpdateName(ID, u.Name)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "이미 사용중인 닉네임입니다",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "성공적으로 변경되었습니다.",
	})
}
