package handler

import (
	"encoding/json"
	"net/http"

	"fitness/internal/model"
	"fitness/internal/storage"
)

type UserHandler struct {
	store storage.UserStorage
}

func NewUserHandler(store storage.UserStorage) *UserHandler {
	return &UserHandler{store: store}
}

func (h *UserHandler) UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users := h.store.GetAll()
		json.NewEncoder(w).Encode(users)
	case http.MethodPost:
		var user model.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request format", http.StatusBadRequest)
			return
		}
		if h.store.Exists(user.Email) {
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}
		h.store.Create(user)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	default:
		http.Error(w, "Invalid query method", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) UserHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Path[len("/users/"):]
	user, exists := h.store.Get(email)
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(user)
	case http.MethodPut:
		var updatedUser model.User
		if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
			http.Error(w, "Invalid request format", http.StatusBadRequest)
			return
		}
		h.store.Update(email, updatedUser)
		json.NewEncoder(w).Encode(updatedUser)
	case http.MethodDelete:
		h.store.Delete(email)
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Invalid query method", http.StatusMethodNotAllowed)
	}
}
