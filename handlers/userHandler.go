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

// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user  body     dtos.CreateUserRequest  true  "User data"
// @Success 200  {object}  dtos.CreateUserResponse "Create user response"
// @Failure 400  {object}  helpers.ApiError "Api Error"
// @Router  /api/user/register [post]
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
