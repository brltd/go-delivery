package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/brltd/delivery/dtos"
	"github.com/brltd/delivery/services"
)

type UserHandler struct {
	UserService services.UserService
}

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var request dtos.CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	}

	user, userErr := h.UserService.CreateUser(request)

	if userErr != nil {
		writeResponse(w, userErr.Code, userErr.AsMessage())
		return
	}

	writeResponse(w, http.StatusCreated, user)
}
