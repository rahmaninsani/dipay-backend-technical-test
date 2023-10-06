package exception

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/helper"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/web"
	"net/http"
)

func HTTPErrorHandler(err error, c echo.Context) {
	var httpError *echo.HTTPError
	var response web.Response
	
	if errors.As(err, &httpError) {
		response = helper.ToResponse(httpError.Code, nil, httpError.Message.(string))
	} else {
		response = helper.ToResponse(http.StatusInternalServerError, nil, err.Error())
	}
	
	if err = c.JSON(response.Status, response); err != nil {
		c.Logger().Error(err)
	}
}
