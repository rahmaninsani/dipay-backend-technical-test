package handler

import (
	"github.com/labstack/echo/v4"
)

type EmployeeHandler interface {
	Create(c echo.Context) error
}
