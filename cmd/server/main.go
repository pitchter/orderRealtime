package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/pitchter/orderRealtime/internal/adapters/handlers"
    "github.com/pitchter/orderRealtime/internal/adapters/redis"
    "github.com/pitchter/orderRealtime/internal/adapters/repositories"
    "github.com/pitchter/orderRealtime/internal/usecases"
)

func main() {
    redis.Init()

    menuRepo := repositories.NewMenuRepository()
    menuUsecase := usecases.NewMenuUsecase(menuRepo)
    menuHandler := handlers.NewMenuHandler(menuUsecase)

    orderRepo := repositories.NewOrderRepository()
    orderUsecase := usecases.NewOrderUsecase(orderRepo)
    orderHandler := handlers.NewOrderHandler(orderUsecase)

    app := fiber.New()

    app.Get("/menu", menuHandler.GetMenu)
    app.Post("/order", orderHandler.CreateOrder)

    app.Listen(":3000")
}
