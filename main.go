package main

import (
	"net/http"

	"github.com/black-dev-x/pos-go-api/configs"
	"github.com/black-dev-x/pos-go-api/internal/entity"
	"github.com/black-dev-x/pos-go-api/internal/infra/database"
	"github.com/black-dev-x/pos-go-api/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	configs.LoadConfig()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productHandler := handlers.NewProductHandler(database.NewProductDB(db))
	userHandler := handlers.NewUserHandler(database.NewUserDB(db))

	http.HandleFunc("POST /products", productHandler.CreateProduct)
	http.HandleFunc("POST /users", userHandler.CreateUser)
	http.ListenAndServe(":8000", nil)
}
