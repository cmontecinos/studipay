package handler

import (
	"bigbytes.io/studipay/view"
	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {
	return view.Index().Render(c.Request().Context(), c.Response())
}

func StudentsHandler(c echo.Context) error {
	return view.Students().Render(c.Request().Context(), c.Response())
}

func PaymentHanlder(c echo.Context) error {
	return view.Payments().Render(c.Request().Context(), c.Response())
}
