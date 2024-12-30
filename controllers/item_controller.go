package controllers

import (
	"encoding/json"
	"net/http"
	"project/services"
)

// gerencia os endpoints de itens
type ItemController struct {
	Service *services.ItemService
}

// cria uma nova instância do controlador de itens
func NewItemController(service *services.ItemService) *ItemController {
	return &ItemController{Service: service}
}

// retorna todos os itens
func (c *ItemController) GetAllItems(w http.ResponseWriter, r *http.Request) {
	items := c.Service.GetAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// adiciona um novo item
func (c *ItemController) AddItem(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Name string `json:"name"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil || requestBody.Name == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	item, err := c.Service.Add(requestBody.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}
