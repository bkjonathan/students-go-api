package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/bkjonathan/students-go-api/internal/storage"
	"github.com/bkjonathan/students-go-api/internal/types"
	"github.com/bkjonathan/students-go-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func Create(storage storage.Storage) http.HandlerFunc {
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

		_, err = storage.SaveStudent(student.Name, student.Email, student.Age)
		slog.Info("user created successfully")

		if err != nil {
			response.WriteJSON(w, http.StatusInternalServerError, response.GenerateErrorResponse(fmt.Errorf("failed to save student")))
			return
		}
		response.WriteJSON(w, http.StatusCreated, student)
		// w.Write([]byte("Create student"))
	}
}
