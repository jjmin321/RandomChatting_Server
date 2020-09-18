package main

import "github.com/labstack/echo"

type mainMethod interface{
	main()
}

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	e.Use(middleware.Recover())
	e.GET("/", )
	e.Logger.Fatal(e.Start(":80"))
	e.Logger.Fatal(e.Start(:8080))
}
