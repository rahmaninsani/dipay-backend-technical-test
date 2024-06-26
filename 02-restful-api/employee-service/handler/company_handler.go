package handler

import (
	"github.com/labstack/echo/v4"
)

type CompanyHandler interface {
	Create(c echo.Context) error
	FindAll(c echo.Context) error
	SetActive(c echo.Context) error
}
