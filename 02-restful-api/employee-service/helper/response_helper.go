package helper

import (
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/web"
	"strconv"
)

func ToResponse(code int, data any, message string) web.Response {
	return web.Response{
		Status:  code,
		Code:    strconv.Itoa(code),
		Data:    data,
		Message: message,
	}
}

func ToAdminLoginResponse(token string) web.AdminLoginResponse {
	return web.AdminLoginResponse{
		Token: token,
	}
}
