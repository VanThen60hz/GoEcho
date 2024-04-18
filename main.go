package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()

	server.GET("/", hello)

	server.POST("/login", login)

	server.Logger.Fatal(server.Start(":8080"))
}

func hello(c echo.Context) error {
	x := &X{
		Text: "Hello World",
	}
	return c.JSON(http.StatusOK, x)
}

func login(c echo.Context) error {
	req := new(LoginRequest)
	c.Bind(req)

	log.Printf("req data %+v", req)

	if req.Username != "admin" || req.Password != "123456" {
		return c.JSON(http.StatusUnauthorized, "username/password invalid!")
	}

	return c.JSON(http.StatusOK, &LoginResponse{
		Token: "123456",
	})
}

type LoginRequest struct {
	Username string `json:"username" form:"username" xml:"username" query:"username"`
	Password string `json:"password" form:"password" xml:"password" query:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type X struct {
	Text string `json:"text"`
}
