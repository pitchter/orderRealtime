package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/pitchter/orderRealtime/internal/adapters/database"
    "github.com/pitchter/orderRealtime/internal/adapters/handlers"
    redisAdapter "github.com/pitchter/orderRealtime/internal/adapters/redis"
    "github.com/pitchter/orderRealtime/internal/adapters/repositories"
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
    orderUsecase := usecases.NewOrderUsecase(orderRepo)
    orderHandler := handlers.NewOrderHandler(orderUsecase)

    app := fiber.New()

    app.Get("/menu", menuHandler.GetMenu)
    app.Post("/menu", menuHandler.CreateMenuItem)
    app.Post("/order", orderHandler.CreateOrder)

    // Initialize and start the order subscriber
    orderSubscriber := subscriber.NewOrderSubscriber(redisClient)
    go orderSubscriber.Subscribe()

    app.Listen(":3000")
}
