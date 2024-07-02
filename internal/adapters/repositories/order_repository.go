package repositories

import (
    "github.com/pitchter/orderRealtime/internal/entities"
)

type OrderRepository interface {
    CreateOrder(order entities.Order) (entities.Order, error)
}

type orderRepository struct {
    orders []entities.Order
}

func NewOrderRepository() OrderRepository {
    return &orderRepository{
        orders: []entities.Order{},
    }
}

func (repo *orderRepository) CreateOrder(order entities.Order) (entities.Order, error) {
    order.ID = len(repo.orders) + 1
    repo.orders = append(repo.orders, order)
    return order, nil
}