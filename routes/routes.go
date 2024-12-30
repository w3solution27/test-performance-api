package routes

import (
	"net/http"
	"project/controllers"
	"project/services"

	"github.com/gorilla/mux"
)

//registra todas as rotas da aplicação
func RegisterRoutes(router *mux.Router) {
	// Instanciar serviço/controlador
	itemService := services.NewItemService()
	itemController := controllers.NewItemController(itemService)

	// rotas
	router.HandleFunc("/items", itemController.GetAllItems).Methods(http.MethodGet)
	router.HandleFunc("/items", itemController.AddItem).Methods(http.MethodPost)
}
