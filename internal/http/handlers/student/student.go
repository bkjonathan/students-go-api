package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/bkjonathan/students-go-api/internal/types"
	"github.com/bkjonathan/students-go-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJSON(w, http.StatusBadRequest, response.GenerateErrorResponse(fmt.Errorf("empty body")))
			return
		}
		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GenerateErrorResponse(fmt.Errorf("invalid request payload")))
			return
		}

		if err := validator.New().Struct(student); err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.ValidationErrorResponse(err.(validator.ValidationErrors)))
			return
		}

		response.WriteJSON(w, http.StatusCreated, student)
		// w.Write([]byte("Create student"))
	}
}
