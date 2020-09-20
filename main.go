package main

import (
	"RandomChatting_Server/config"
	"RandomChatting_Server/controller/chattingserver"
	"RandomChatting_Server/database"

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
	e.Logger.Fatal(e.Start(":80"))
}
