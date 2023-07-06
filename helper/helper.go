package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func APIresponse(status int, data interface{}) Response {
	response := Response{
		Status: status,
		Data:   data,
	}
	return response
}

type ResponseEmail struct {
	Status int         `json:"status"`
	Message string `json:"message"`
	Data   interface{} `json:"data"`
}

func APIresponseEmail(message string , status int, data interface{}) ResponseEmail {
	responseEmail := ResponseEmail{
		Status: status,
		Message: message,
		Data:   data,
	}
	return responseEmail
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
