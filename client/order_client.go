package main

import (
	"context"
	"fmt"
	"github.com/GaleDetail/go-microservices/api/order"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// Підключаємося до сервера
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Не вдалося підключитися: %v", err)
	}
	defer conn.Close()

	client := order.NewOrderServiceClient(conn)

	// Встановлюємо таймаут для запиту
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Виконуємо запит GetOrder
	res, err := client.GetOrder(ctx, &order.OrderRequest{Id: "456"})
	if err != nil {
		log.Fatalf("Не вдалося отримати замовлення: %v", err)
	}

	fmt.Printf("Отримано замовлення: ID=%s, Description=%s\n", res.Id, res.Description)
}
