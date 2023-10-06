package handler

import (
	"github.com/labstack/echo/v4"
)

type AdminHandler interface {
	Login(c echo.Context) error
}
