package controller

import (
	"RandomChatting_Server/lib"
	"RandomChatting_Server/model"

	"github.com/labstack/echo"
)

type signInMethod interface {
	SignIn()
}

// SignInParam - 파라미터 형식 정의 구조체
type SignInParam struct {
	ID string `json:"id" form:"id" query:"id"`
	Pw string `json:"pw" form:"pw" query:"pw"`
}

// SignIn - 로그인 메서드
func SignIn(c echo.Context) error {
	u := new(SignInParam)
	if err := c.Bind(u); err != nil {
		return err
	}
	_, err := model.FindMember(u.ID, u.Pw)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"status":  400,
			"message": "해당 정보에 맞는 유저가 없습니다",
		})
	}
	accessToken, err := lib.CreateAccessToken(u.ID, u.Pw)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "토큰 생성 중 오류",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":      200,
		"message":     "로그인 성공",
		"accessToken": accessToken,
	})
}
