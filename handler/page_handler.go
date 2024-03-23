package handler

import (
	"bigbytes.io/studipay/service"
	"bigbytes.io/studipay/view/page"
	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {
	return page.Index().Render(c.Request().Context(), c.Response())
}

func StudentsHandler(c echo.Context) error {
	students, err := service.GetAllStudents()
	if err != nil {
		return c.String(500, "Internal Server Error")
	}
	return page.Students(students).Render(c.Request().Context(), c.Response())
}

func PaymentHanlder(c echo.Context) error {
	return page.Payments().Render(c.Request().Context(), c.Response())
}
