package main

import (
	"net/http"

	"github.com/black-dev-x/pos-go-api/configs"
	"github.com/black-dev-x/pos-go-api/internal/entity"
	"github.com/black-dev-x/pos-go-api/internal/infra/database"
	"github.com/black-dev-x/pos-go-api/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
func main() {
	configs.LoadConfig()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productHandler := handlers.NewProductHandler(database.NewProductDB(db))
	userHandler := handlers.NewUserHandler(database.NewUserDB(db))

	router := chi.NewRouter()
	router.Use(jsonContentTypeMiddleware)
	router.Post("/products", productHandler.CreateProduct)
	router.Put("/products/{id}", productHandler.UpdateProduct)
	router.Get("/products", productHandler.GetProducts)
	router.Get("/products/{id}", productHandler.GetProductByID)

	router.Post("/users", userHandler.CreateUser)
	// http.Handle("/", router)

	// http.HandleFunc("POST /products", productHandler.CreateProduct)
	// http.HandleFunc("POST /users", userHandler.CreateUser)
	http.ListenAndServe(":8000", router)
}
