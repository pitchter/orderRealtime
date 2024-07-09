package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/pitchter/orderRealtime/internal/adapters/database"
	"github.com/pitchter/orderRealtime/internal/adapters/handlers"
	redisAdapter "github.com/pitchter/orderRealtime/internal/adapters/redis"
	"github.com/pitchter/orderRealtime/internal/adapters/repositories"
	"github.com/pitchter/orderRealtime/internal/adapters/websockets"
	"github.com/pitchter/orderRealtime/internal/subscriber"
	"github.com/pitchter/orderRealtime/internal/usecases"
)

func main() {
	database.Init()
	redisClient := redisAdapter.Init()

	menuRepo := repositories.NewMenuRepository(database.DB)
	menuUsecase := usecases.NewMenuUsecase(menuRepo)
	menuHandler := handlers.NewMenuHandler(menuUsecase)

	orderRepo := repositories.NewOrderRepository(database.DB)
	orderUsecase := usecases.NewOrderUsecase(orderRepo, menuRepo)
	orderHandler := handlers.NewOrderHandler(orderUsecase)

	app := fiber.New()

	app.Get("/menu", menuHandler.GetMenu)
	app.Post("/menu", menuHandler.CreateMenuItem)
	app.Post("/order", orderHandler.CreateOrder)

	// Initialize and start the WebSocket handler
	wsHandler := websockets.NewWebSocketHandler()
	go wsHandler.HandleMessages()

	// Initialize and start the order subscriber
	orderSubscriber := subscriber.NewOrderSubscriber(redisClient, wsHandler)
	go orderSubscriber.Subscribe()

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		wsHandler.HandleConnections(c)
	}))

	app.Listen(":3000")
}
