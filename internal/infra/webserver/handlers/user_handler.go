package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/black-dev-x/pos-go-api/internal/dto"
	"github.com/black-dev-x/pos-go-api/internal/entity"
	"github.com/black-dev-x/pos-go-api/internal/infra/database"
)

type UserHandler struct {
	UserDB database.UserDBInterface
}

func NewUserHandler(db database.UserDBInterface) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserDTO dto.CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&createUserDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := entity.NewUser(createUserDTO.Email, createUserDTO.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.UserDB.Create(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
