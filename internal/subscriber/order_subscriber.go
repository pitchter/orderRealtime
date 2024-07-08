package subscriber

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "github.com/pitchter/orderRealtime/internal/entities"
    "github.com/redis/go-redis/v9"
)

// OrderSubscriber handles subscription to order-related events.
type OrderSubscriber struct {
    redisClient *redis.Client
}

// NewOrderSubscriber creates a new OrderSubscriber.
func NewOrderSubscriber(redisClient *redis.Client) *OrderSubscriber {
    return &OrderSubscriber{redisClient: redisClient}
}

// Subscribe subscribes to the order created events.
func (os *OrderSubscriber) Subscribe() {
    ctx := context.Background()
    pubsub := os.redisClient.Subscribe(ctx, "order_created")
    ch := pubsub.Channel()

    for msg := range ch {
        var order entities.Order
        err := json.Unmarshal([]byte(msg.Payload), &order)
        if err != nil {
            log.Printf("Error unmarshalling order: %v", err)
            continue
        }
        fmt.Printf("Order Created: %+v\n", order)
    }
}
