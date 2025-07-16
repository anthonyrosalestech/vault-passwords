// handlers/handler.go
package handlers

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/register", h.RegisterUser).Methods("POST")
	r.HandleFunc("/login", h.LoginUser).Methods("POST")
	r.HandleFunc("/vault", h.CreateVaultEntry).Methods("POST")
	r.HandleFunc("/vault", h.GetVaultEntries).Methods("GET")
	// Agrega más rutas según sea necesario
}
