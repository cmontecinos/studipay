package handler

import (
	"log"
	"net/http"

	"bigbytes.io/studipay/persistence"
	"bigbytes.io/studipay/service"
	"bigbytes.io/studipay/types"
	"github.com/labstack/echo/v4"
)

func CreateStudentHandler(c echo.Context) error {
	student := new(types.Student)
	if err := c.Bind(student); err != nil {
		log.Println("Error binding student:", err)
		return c.String(http.StatusBadRequest, "bad request")
	}

	created, err := service.CreateStudent(student)
	if err != nil {
		log.Println("Error creating student:", err)

		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, created)
}

func GetStudentHandler(c echo.Context) error {
	rut := c.Param("rut")
	student, err := service.GetStudent(rut)
	if err != nil {
		if err == persistence.ErrStudentNotFound {
			return c.String(http.StatusNotFound, "not found")
		}
		log.Println("Error getting student:", err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, student)
}

func UpdateStudentHandler(c echo.Context) error {
	student := new(types.Student)
	if err := c.Bind(student); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	err := service.UpdateStudent(student)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	c.Response().Writer.Write([]byte("updated"))
	return nil
}
