package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/laurati/api_product_user/configs"
	"github.com/laurati/api_product_user/internal/entity"
	"github.com/laurati/api_product_user/internal/infra/database"
	"github.com/laurati/api_product_user/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()

	// r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)
	// r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	// r.Use(middleware.WithValue("JwtExperesIn", configs.JwtExperesIn))

	r.Route("/products", func(r chi.Router) {
		// r.Use(jwtauth.Verifier(configs.TokenAuth))
		// r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		// r.Get("/", productHandler.GetProducts)
		// r.Get("/{id}", productHandler.GetProduct)
		// r.Put("/{id}", productHandler.UpdateProduct)
		// r.Delete("/{id}", productHandler.DeleteProduct)
	})

	http.ListenAndServe(":8000", r)
}
