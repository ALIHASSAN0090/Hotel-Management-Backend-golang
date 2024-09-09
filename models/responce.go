package models

type ErrorResponse struct {
	Error string `json:"error"`
}

type Response struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}
