package handler

import (
	"net/http"

	"bigbytes.io/studipay/service"
	"bigbytes.io/studipay/types"
	"github.com/labstack/echo/v4"
)

func CreatePaymentHandler(c echo.Context) error {
	payment := new(types.Payment)
	if err := c.Bind(payment); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	err := service.CreatePayment(payment)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	c.Response().Writer.Write([]byte("created"))
	return nil

}

func GetPaymentByStudentIdHandler(c echo.Context) error {
	studentID := c.Param("studentID")
	payments, err := service.GetPaymentsByStudentID(studentID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, payments)
}
