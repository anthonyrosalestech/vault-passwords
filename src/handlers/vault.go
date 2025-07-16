// handlers/vault.go
package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"vault_backend/models"
	"github.com/google/uuid"
)

type VaultEntryRequest struct {
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	WebsiteURL string `json:"website_url"`
	Notes    string `json:"notes"`
}

func (h *Handler) CreateVaultEntry(w http.ResponseWriter, r *http.Request) {
	var req VaultEntryRequest
	_ = json.NewDecoder(r.Body).Decode(&req)

	entry := models.VaultEntry{
		ID: uuid.New().String(),
		UserID: req.UserID,
		Name: req.Name,
		Username: req.Username,
		PasswordEnc: req.Password, // Aquí puedes aplicar cifrado con AES
		WebsiteURL: req.WebsiteURL,
		Notes: req.Notes,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	h.DB.Create(&entry)
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetVaultEntries(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	var entries []models.VaultEntry
	h.DB.Where("user_id = ?", userID).Find(&entries)
	json.NewEncoder(w).Encode(entries)
}