package usecases

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strconv"

	"github.com/pitchter/orderRealtime/internal/adapters/repositories"
	"github.com/pitchter/orderRealtime/internal/entities"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type OrderUsecase struct {
	orderRepo   repositories.OrderRepository
	menuRepo    repositories.MenuRepository
	redisClient *redis.Client
}

func NewOrderUsecase(orderRepo repositories.OrderRepository, menuRepo repositories.MenuRepository, redisClient *redis.Client) *OrderUsecase {
	return &OrderUsecase{
		orderRepo:   orderRepo,
		menuRepo:    menuRepo,
		redisClient: redisClient,
	}
}

func (uc *OrderUsecase) CreateOrder(order entities.Order) (entities.Order, error) {
	var total float64
	var fullItems []entities.OrderItem

	for _, item := range order.Items {
		menuItem, err := uc.menuRepo.GetMenuItemByID(item.MenuItemID)
		if err != nil {
			return order, errors.New("menu item not found")
		}
		total += menuItem.Price * float64(item.Quantity)
		fullItems = append(fullItems, item)
	}

	order.Total = total
	order.Items = fullItems

	// Save the order in the database
	createdOrder, err := uc.orderRepo.CreateOrder(order)
	if err != nil {
		return createdOrder, err
	}

	// Convert order ID from uint to string using strconv.Itoa
	orderIDStr := strconv.Itoa(int(createdOrder.ID))
	redisKey := "order:" + orderIDStr

	// Save the order in Redis
	orderJSON, err := json.Marshal(createdOrder)
	if err != nil {
		log.Printf("Failed to marshal order: %v", err)
		return createdOrder, err
	}

	err = uc.redisClient.Set(ctx, redisKey, orderJSON, 0).Err()
	if err != nil {
		log.Printf("Failed to save order in Redis: %v", err)
		return createdOrder, err
	}

	return createdOrder, nil
}

func (uc *OrderUsecase) Reorder(orderID uint) (entities.Order, error) {
	// Convert order ID to string for Redis lookup
	orderIDStr := strconv.Itoa(int(orderID))
	redisKey := "order:" + orderIDStr

	// Attempt to get the order from Redis
	val, err := uc.redisClient.Get(ctx, redisKey).Result()
	if err == redis.Nil {
		return entities.Order{}, errors.New("Order not found")
	} else if err != nil {
		log.Printf("Failed to get order from Redis: %v", err)
		return entities.Order{}, err
	}

	var oldOrder entities.Order
	if err := json.Unmarshal([]byte(val), &oldOrder); err != nil {
		log.Printf("Failed to unmarshal order: %v", err)
		return entities.Order{}, err
	}

	// Create a new order with the same items and customer info
	newOrder := entities.Order{
		CustomerName: oldOrder.CustomerName,
		TableNumber:  oldOrder.TableNumber,
		Items:        oldOrder.Items,
		Status:       "pending", // Reset status for the new order
	}

	// Save the new order in the database
	createdOrder, err := uc.CreateOrder(newOrder)
	if err != nil {
		return createdOrder, err
	}

	return createdOrder, nil
}
