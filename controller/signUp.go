package controller

import (
	"RandomChatting_Server/model"

	"github.com/labstack/echo"
)

type signUpMethod interface {
	SignUp()
}

// SignUpParam - 파라미터 형식 정의 구조체
type SignUpParam struct {
	ID   string `json:"id" form:"id" query:"id"`
	Pw   string `json:"pw" form:"pw" query:"pw"`
	Name string `json:"name" form:"name" query:"name"`
}

// SignUp - 회원가입 API
func SignUp(c echo.Context) error {
	u := new(SignUpParam)
	if err := c.Bind(u); err != nil {
		return err
	}
	if u.ID == "" || u.Name == "" || u.Pw == "" {
		return c.JSON(400, map[string]interface{}{
			"status":  400,
			"message": "모든 값을 입력해주세요",
		})
	}
	err := model.CheckDupID(u.ID)
	if err == nil {
		return c.JSON(400, map[string]interface{}{
			"status":  400,
			"message": "이미 사용중인 아이디입니다",
		})
	}
	err = model.CheckDupName(u.Name)
	if err == nil {
		return c.JSON(400, map[string]interface{}{
			"status":  400,
			"message": "이미 사용중인 닉네임입니다",
		})
	}
	err = model.CreateMember(u.ID, u.Name, u.Pw)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "멤버 생성 중 오류 발생",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"message": "멤버 생성 완료",
	})
}
