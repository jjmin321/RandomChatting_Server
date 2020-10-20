package lib

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type jwtMethod interface {
	CreateAccessToken()
	VerifyAccessToken()
}

// CreateAccessToken : 액세스 토큰 생성
func CreateAccessToken(ID, Pw string) (string, error) {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	claims := accessToken.Claims.(jwt.MapClaims)
	claims["ID"] = ID
	claims["Pw"] = Pw
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	t, err := accessToken.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

// VerifyAccessToken : 액세스 토큰 검증
func VerifyAccessToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		ID := claims["ID"].(string)
		Pw := claims["Pw"].(string)
		c.Set("ID", ID)
		c.Set("Pw", Pw)
		return next(c)
	}
}
