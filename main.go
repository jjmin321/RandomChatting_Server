package main

import (
	"RandomChatting_Server/config"
	"RandomChatting_Server/controller"
	"RandomChatting_Server/controller/chattingserver"
	"RandomChatting_Server/database"
	"RandomChatting_Server/lib/jwt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type mainMethod interface {
	main()
}

func main() {
	config.InitConfig()
	database.Connect()
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	e.Use(middleware.Recover())
	e.GET("/chatting", chattingserver.Socket)
	e.POST("/signIn", controller.SignIn)
	e.POST("/signUp", controller.SignUp)
	e.PUT("/putImage", controller.PutImage, middleware.JWT([]byte("secret")), jwt.VerifyAccessToken)
	e.Logger.Fatal(e.Start(":80"))
}
