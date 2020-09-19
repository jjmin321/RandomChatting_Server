package main

import (
	socket "RandomChatting_Server/controller/chatting/Socket"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type mainMethod interface {
	main()
}

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	e.Use(middleware.Recover())
	e.GET("/", socket.Socket)
	e.Logger.Fatal(e.Start(":80"))
}
