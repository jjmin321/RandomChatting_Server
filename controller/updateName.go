package controller

import (
	"RandomChatting_Server/database"

	"github.com/labstack/echo"
)

// UpdateNameParam - 파라미터 형식 정의 구조체
type UpdateNameParam struct {
	Name string `json:"name" form:"name" query:"name"`
}

// UpdateMyInfo - 내 정보 변경 메서드
func UpdateMyInfo(c echo.Context) error {
	ID := c.Get("ID")
	u := new(UpdateNameParam)
	if err := c.Bind(u); err != nil {
		return err
	}
	Member := &database.Member{}
	err := database.DB.Model(Member).Where("id = ?", ID).Update("name", u.Name).Error
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
