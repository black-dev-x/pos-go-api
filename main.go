package main

import (
	"net/http"

	"github.com/black-dev-x/pos-go-api/configs"
	"github.com/black-dev-x/pos-go-api/internal/entity"
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

	http.ListenAndServe(":8080", nil)
}
