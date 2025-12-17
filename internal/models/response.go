package models

type ResponseError struct {
	Error string `json:"error"`
}

type Response struct {
	Data any `json:"data"`
}
