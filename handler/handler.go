package handler

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandlerRepository(db *gorm.DB) *Handler {
	return &Handler{
		DB: db,
	}
}

func Response(writer http.ResponseWriter, code int, body any) {
	b, err := json.Marshal(body)

	if err != nil {
		ErrorResponse(writer, 500, err.Error())
		return
	}

	writer.WriteHeader(code)
	_, _ = writer.Write(b)
}

func ErrorResponse(writer http.ResponseWriter, code int, message string) {
	writer.WriteHeader(code)

	response := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  "error",
		Message: message,
	}

	b, _ := json.Marshal(response)
	_, _ = writer.Write(b)
}

func DBError(tx *gorm.DB, w http.ResponseWriter) bool {
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			ErrorResponse(w, 404, tx.Error.Error())
		} else if errors.Is(tx.Error, gorm.ErrDuplicatedKey) {
			ErrorResponse(w, 422, "Duplicate value!")
		} else {
			ErrorResponse(w, 500, "Uh oh! Something went wrong")
		}

		return true
	}

	return false
}
