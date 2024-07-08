package usecases

import (
    "github.com/pitchter/orderRealtime/internal/entities"
    "github.com/pitchter/orderRealtime/internal/adapters/repositories"
    "github.com/pitchter/orderRealtime/internal/service"
)

type OrderUsecase struct {
    orderRepo repositories.OrderRepository
}

func NewOrderUsecase(repo repositories.OrderRepository) *OrderUsecase {
    return &OrderUsecase{orderRepo: repo}
}

func (uc *OrderUsecase) CreateOrder(order entities.Order) (entities.Order, error) {
    createdOrder, err := uc.orderRepo.CreateOrder(order)
    if err != nil {
        return createdOrder, err
    }
    // Publish order created event
    err = services.PublishOrderCreated(createdOrder)
    if err != nil {
        return createdOrder, err
    }
    return createdOrder, nil
}
