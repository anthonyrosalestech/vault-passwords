// handlers/user.go
package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"vault_backend/models"
	"vault_backend/utils"
	"github.com/google/uuid"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	_ = json.NewDecoder(r.Body).Decode(&req)

	salt := utils.GenerateSalt()
	hash := utils.HashPassword(req.Password, salt)

	user := models.User{
		ID:           uuid.New().String(),
		Email:        req.Email,
		Username:     req.Username,
		PasswordHash: hash,
		PasswordSalt: salt,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	h.DB.Create(&user)
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	_ = json.NewDecoder(r.Body).Decode(&req)

	var user models.User
	h.DB.Where("email = ?", req.Email).First(&user)
	if utils.HashPassword(req.Password, user.PasswordSalt) != user.PasswordHash {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}