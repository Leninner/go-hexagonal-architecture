package core

import "github.com/go-playground/validator/v10"

type ResponseMessage struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	FailedField string `json:"failedField"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func NewResponseMessage(message string, data interface{}) ResponseMessage {
	return ResponseMessage{
		Message: message,
		Data:    data,
	}
}

func ValidateStruct(s interface{}) []*ErrorResponse {

	var errors []*ErrorResponse
	validate := validator.New()

	err := validate.Struct(s)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			error := &ErrorResponse{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			}

			errors = append(errors, error)
		}
	}

	return errors
}
