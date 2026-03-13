package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusSuccess = "success"
	StatusError   = "error"
)

func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		return json.NewEncoder(w).Encode(data)
	}
	return nil
}

func GenerateErrorResponse(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationErrorResponse(err validator.ValidationErrors) Response {
	var errorMessages []string
	for _, e := range err {
		switch e.ActualTag() {
		case "required":
			errorMessages = append(errorMessages, fmt.Sprintf("%s is required", e.Field()))
		case "gte":
			errorMessages = append(errorMessages, fmt.Sprintf("%s must be greater than or equal to %s", e.Field(), e.Param()))
		case "lte":
			errorMessages = append(errorMessages, fmt.Sprintf("%s must be less than or equal to %s", e.Field(), e.Param()))
		default:
			errorMessages = append(errorMessages, fmt.Sprintf("%s is invalid", e.Field()))
		}
	}
	return Response{
		Status: StatusError,
		Error:  strings.Join(errorMessages, ", "),
		// Error:  "validation error: " + fmt.Sprintf("%v", errorMessages),
	}
}
