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
	//page
	e.GET("/students", handler.StudentsHandler)
	e.GET("/payments", handler.PaymentHanlder)

	//CRUD operations
	//students
	e.POST("/students", handler.CreateStudentHandler)
	e.GET("/students/:rut", handler.GetStudentHandler)
	e.PUT("/students", handler.UpdateStudentHandler)
	e.GET("/students/search", handler.SearchStudentByNameHandler)

	//payments
	e.POST("/payments", handler.CreatePaymentHandler)
	e.GET("/payments/:studentID", handler.GetPaymentByStudentIdHandler)

	e.Start(":8080")
}
