package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tharun13055/golang/mux/services"
)
// step 1
type UserHandler struct {
	Service services.UserService
}

// step 3 it return to route
func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

// step 2
func (u *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := u.Service.GetUserById(id)
	if err != nil {
		http.Error(w, "There is no user id", http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
