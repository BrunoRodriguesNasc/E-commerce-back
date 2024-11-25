package controllers

import (
	"encoding/json"
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
	"net/http"
)

type ProductController struct {
	repo *repositories.ProductRepository
}

func NewProductController(repo *repositories.ProductRepository) *ProductController {
	return &ProductController{repo: repo}
}

func (c *ProductController) Create(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.repo.Create(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func (c *ProductController) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := c.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
