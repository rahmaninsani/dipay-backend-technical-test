package exception

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/web"
	"net/http"
)

func HTTPErrorHandler(err error, c echo.Context) {
	var httpError *echo.HTTPError
	var response *web.Response
	
	if errors.As(err, &httpError) {
		response = &web.Response{
			Status:  httpError.Code,
			Code:    string(rune(httpError.Code)),
			Data:    nil,
			Message: httpError.Message.(string),
		}
	} else {
		response = &web.Response{
			Status:  http.StatusInternalServerError,
			Code:    string(rune(http.StatusInternalServerError)),
			Data:    nil,
			Message: err.Error(),
		}
	}
	
	if err = c.JSON(response.Status, response); err != nil {
		c.Logger().Error(err)
	}
}
