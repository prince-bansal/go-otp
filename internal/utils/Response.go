package Response

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field string
	Error string
}

type ErrorResponse struct {
	Status      int         `json:"status"`
	Error       string      `json:"error"`
	ErrorDetail interface{} `json:"errorDetail"`
}

func (e *ErrorResponse) SendInvalidError(message string, details interface{}) *ErrorResponse {
	return &ErrorResponse{
		Status:      400,
		Error:       message,
		ErrorDetail: details,
	}
}

func SendValidationError(err error) *ErrorResponse {
	return &ErrorResponse{
		Status:      400,
		Error:       "validation Error",
		ErrorDetail: getError(err),
	}
}

func SendError(message string, details interface{}) *ErrorResponse {
	return &ErrorResponse{
		Status:      400,
		Error:       message,
		ErrorDetail: details,
	}
}

func getError(err error) interface{} {
	var details []string

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			errorMessage := fmt.Sprintf("%s %s", fe.Field(), generateMessage(fe))
			details = append(details, errorMessage)
		}
		return details
	} else {
		return err
	}
}

func generateMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "is required"
	case "email":
		return "must be a valid email"
	case "min":
		return fmt.Sprintf("must be at least %s characters", fe.Param())
	case "max":
		return fmt.Sprintf("must be at most %s characters", fe.Param())
	default:
		return fmt.Sprintf("failed '%s' validation", fe.Tag())
	}
}
