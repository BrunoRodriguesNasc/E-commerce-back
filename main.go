package main

import (
	"e-commerce/internal/routes"
	"e-commerce/pkg/database"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Inicializa conexão com banco
	db := database.Connect()
	defer database.Close(db)

	// Setup rotas
	router := routes.SetupRoutes()

	fmt.Println("Server starting on port 3000...")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}