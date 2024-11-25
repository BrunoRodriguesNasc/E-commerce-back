package routes

import (
	"e-commerce/internal/controllers"
	"e-commerce/internal/repositories"
	"e-commerce/pkg/database"
	"net/http"
)

func SetupRoutes() http.Handler {
	db := database.Connect()
	
	productRepo := repositories.NewProductRepository(db)
	productController := controllers.NewProductController(productRepo)

	mux := http.NewServeMux()
	
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productController.GetAll(w, r)
		case http.MethodPost:
			productController.Create(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
