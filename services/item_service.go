package services

import (
	"errors"
	"project/models"
)

// gerencia a lógica de negócio para itens
type ItemService struct {
	items []models.Item
}

// cria uma nova instância do serviço de itens
func NewItemService() *ItemService {
	return &ItemService{
		items: []models.Item{},
	}
}

// retorna todos os itens
func (s *ItemService) GetAll() []models.Item {
	return s.items
}

// adiciona um novo item
func (s *ItemService) Add(name string) (models.Item, error) {
	if name == "" {
		return models.Item{}, errors.New("name is required")
	}
	item := models.Item{
		ID:   len(s.items) + 1,
		Name: name,
	}
	s.items = append(s.items, item)
	return item, nil
}
