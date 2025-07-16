// main.go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"vault_backend/handlers"
	"vault_backend/models"
)

func main() {
	dsn := "host=localhost user=postgres password=yourpass dbname=vault port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	db.AutoMigrate(
		&models.User{},
		&models.VaultEntry{},
		&models.Tag{},
		&models.VaultEntryTag{},
		&models.VaultEntryHistory{},
		&models.UserDevice{},
		&models.WebsiteCatalog{},
		&models.WebsiteDomain{},
		&models.WebsiteCategory{},
	)

	r := mux.NewRouter()
	h := handlers.NewHandler(db)

	h.RegisterRoutes(r)

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", r)
}