package subscriber

import (
	"context"
	"encoding/json"
	"log"

	"github.com/pitchter/orderRealtime/internal/adapters/websockets"
	"github.com/pitchter/orderRealtime/internal/entities"
	"github.com/redis/go-redis/v9"
)

type OrderSubscriber struct {
	redisClient *redis.Client
	wsHandler   *websockets.WebSocketHandler
}

func NewOrderSubscriber(redisClient *redis.Client, wsHandler *websockets.WebSocketHandler) *OrderSubscriber {
	return &OrderSubscriber{
		redisClient: redisClient,
		wsHandler:   wsHandler,
	}
}

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
		os.wsHandler.Broadcast <- order
	}
}
