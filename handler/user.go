package handler

import (
	"dockerHomework/model"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h *Handler) Users(r chi.Router) {
	r.Get("/{id:[1-9][0-9]*}", func(writer http.ResponseWriter, request *http.Request) {
		ID, _ := strconv.Atoi(chi.URLParam(request, "id"))

		user := model.User{}
		tx := h.DB.First(&user, ID)

		if DBError(tx, writer) {
			return
		}

		Response(writer, 200, user)
	})

	r.Post("/", func(writer http.ResponseWriter, request *http.Request) {
		user := model.User{}

		//@todo Make some field validation in future
		err := json.NewDecoder(request.Body).Decode(&user)
		if err != nil {
			ErrorResponse(writer, 400, err.Error())
			return
		}

		tx := h.DB.Create(&user)
		if DBError(tx, writer) {
			return
		}

		Response(writer, 201, &user)
	})

	r.Put("/{id:[1-9][0-9]*}", func(writer http.ResponseWriter, request *http.Request) {
		ID, _ := strconv.Atoi(chi.URLParam(request, "id"))
		user := model.User{}

		tx := h.DB.First(&model.User{}, ID)

		if DBError(tx, writer) {
			return
		}

		//@todo Make some field validation in future
		err := json.NewDecoder(request.Body).Decode(&user)
		if err != nil {
			ErrorResponse(writer, 400, err.Error())
			return
		}
		user.ID = uint(ID)

		tx = h.DB.Save(&user)
		if DBError(tx, writer) {
			return
		}

		Response(writer, 200, &user)
	})

	r.Delete("/{id:[1-9][0-9]*}", func(writer http.ResponseWriter, request *http.Request) {
		ID, _ := strconv.Atoi(chi.URLParam(request, "id"))
		h.DB.Delete(model.User{}, ID)

		Response(writer, 204, nil)
	})
}
