package main

import (
	"fmt"

	"bigbytes.io/studipay/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Starting server...")
	e := echo.New()
	e.Static("/static", "static")
	e.GET("/", handler.IndexHandler)

	e.GET("/students", handler.StudentsHandler)

	e.Start(":8080")
}
