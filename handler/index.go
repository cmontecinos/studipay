package handler

import (
	"bigbytes.io/studipay/view"
	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {
	return view.Index().Render(c.Request().Context(), c.Response())
}
