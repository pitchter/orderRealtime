package services

import (
    "context"
    "github.com/pitchter/orderRealtime/internal/adapters/redis"
    "github.com/pitchter/orderRealtime/internal/entities"
    "encoding/json"
)

// PublishOrderCreated publishes an event to the "order_created" channel when a new order is created.
func PublishOrderCreated(order entities.Order) error {
    ctx := context.Background()
    // Serialize the order to JSON
    orderJson, err := json.Marshal(order)
    if err != nil {
        return err
    }
    // Publish the order JSON to the "order_created" channel
    return redis.Rdb.Publish(ctx, "order_created", orderJson).Err()
}
