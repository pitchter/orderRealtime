package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pitchter/orderRealtime/internal/entities"
	"github.com/pitchter/orderRealtime/internal/usecases"
	"github.com/pitchter/orderRealtime/internal/utils"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type OrderHandler struct {
	orderUsecase *usecases.OrderUsecase
	redisClient  *redis.Client
}

func NewOrderHandler(orderUsecase *usecases.OrderUsecase, redisClient *redis.Client) *OrderHandler {
	return &OrderHandler{
		orderUsecase: orderUsecase,
		redisClient:  redisClient,
	}
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var order entities.Order
	if err := c.BodyParser(&order); err != nil {
		return utils.HandleError(c, err)
	}

	if order.CustomerName == "" {
		return utils.HandleError(c, errors.New("customer name is required"))
	}

	// ดึงเลขโต๊ะจาก query parameter
	//  tableNumber, err := c.QueryInt("table", 0)
	//  if err != nil || tableNumber <= 0 {
	// 	 return utils.HandleError(c, errors.New("invalid table number"))
	//  }

	// ตั้งค่า table number ในคำสั่งซื้อ
	//  order.TableNumber = tableNumber

	if order.TableNumber <= 0 {
		return utils.HandleError(c, errors.New("invalid table number"))
	}

	for _, item := range order.Items {
		if item.Quantity <= 0 {
			return utils.HandleError(c, errors.New("quantity must be greater than 0"))
		}
	}

	createdOrder, err := h.orderUsecase.CreateOrder(order)
	if err != nil {
		return utils.HandleError(c, err)
	}

	return c.JSON(createdOrder)
}

func (h *OrderHandler) GetOrder(c *fiber.Ctx) error {
	orderID := c.Params("id")

	// Attempt to get the order from Redis
	val, err := h.redisClient.Get(ctx, "order:"+orderID).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Order not found"})
	} else if err != nil {
		log.Printf("Failed to get order from Redis: %v", err)
		return utils.HandleError(c, err)
	}

	var order entities.Order
	if err := json.Unmarshal([]byte(val), &order); err != nil {
		log.Printf("Failed to unmarshal order: %v", err)
		return utils.HandleError(c, err)
	}

	return c.JSON(order)
}

func (h *OrderHandler) Reorder(c *fiber.Ctx) error {
	orderIDStr := c.Params("id")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil || orderID <= 0 {
		return utils.HandleError(c, errors.New("Invalid order ID"))
	}

	// Call the usecase to create a new order based on the previous one
	newOrder, err := h.orderUsecase.Reorder(uint(orderID))
	if err != nil {
		return utils.HandleError(c, err)
	}

	return c.JSON(newOrder)
}
