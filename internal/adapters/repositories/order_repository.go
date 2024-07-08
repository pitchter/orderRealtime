package repositories

import (
    "github.com/pitchter/orderRealtime/internal/entities"
    "gorm.io/gorm"
)

type OrderRepository interface {
    CreateOrder(order entities.Order) (entities.Order, error)
}

type orderRepository struct {
    db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
    return &orderRepository{db: db}
}

func (repo *orderRepository) CreateOrder(order entities.Order) (entities.Order, error) {
    result := repo.db.Create(&order)
    return order, result.Error
}
