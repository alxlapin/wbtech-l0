package service

import (
	"wbtech-l0/internal/domain"

	"gorm.io/gorm"
)

// TODO: добавить кеширующий слой в сервис (по ТЗ в задаче кеш должен содержать N последних заказов).
// TODO: мысли по неймингу: вероятно, лучше подойдет *Repository

type OrderService struct {
	GormDB *gorm.DB
}

func (s *OrderService) GetOrderByID(id string) (domain.Order, error) {
	var order domain.Order

	err := s.GormDB.Preload("Delivery").Preload("Payment").Preload("Items").First(&order, "id = ?", id).Error

	return order, err
}
