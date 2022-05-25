package models

type ServerError struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}
