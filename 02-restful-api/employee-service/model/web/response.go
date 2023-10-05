package web

type Response struct {
	Status  int         `json:"status"`
	Code    string      `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
