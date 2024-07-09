package usecases

import (
	"errors"

	"github.com/pitchter/orderRealtime/internal/adapters/repositories"
	"github.com/pitchter/orderRealtime/internal/entities"
	services "github.com/pitchter/orderRealtime/internal/service"
)

type OrderUsecase struct {
	orderRepo repositories.OrderRepository
	menuRepo  repositories.MenuRepository
}

func NewOrderUsecase(orderRepo repositories.OrderRepository, menuRepo repositories.MenuRepository) *OrderUsecase {
	return &OrderUsecase{orderRepo: orderRepo, menuRepo: menuRepo}
}

func (uc *OrderUsecase) CreateOrder(order entities.Order) (entities.Order, error) {
	var total float64

	for _, item := range order.Items {
		menuItem, err := uc.menuRepo.GetMenuItemByID(item.ID)
		if err != nil {
			return order, errors.New("menu item not found")
		}
		total += menuItem.Price
	}

	order.Total = total
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
