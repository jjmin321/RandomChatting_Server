package controller

import (
	"RandomChatting_Server/model"

	"github.com/labstack/echo"
)

type getInfoMethod interface {
	GetInfo()
}

// GetInfo - 유저 정보 읽기 API
func GetInfo(c echo.Context) error {
	Name := c.Get("Name").(string)
	Pw := c.Get("Pw").(string)
	Member, err := model.FindMember(Name, Pw)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "멤버 조회 실패",
			"Member":  nil,
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "멤버 조회 완료",
		"Member":  Member,
	})
}
