package main

import (
	"GoEcho/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	server := echo.New()

	server.Use(middleware.Logger())

	server.GET("/", handler.Hello)
	server.POST("/login", handler.Login, middleware.BasicAuth(middleware.BasicAuth)

	server.Logger.Fatal(server.Start(":8080"))
}

type LoginRequest struct {
	Username string `json:"username" form:"username" xml:"username" query:"username"`
	Password string `json:"password" form:"password" xml:"password" query:"password"`
}
