package main

import (
	"context"
	"fmt"
	"github.com/GaleDetail/go-microservices/api/user"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// Підключаємося до сервера
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Не вдалося підключитися: %v", err)
	}
	defer conn.Close()

	client := user.NewUserServiceClient(conn)

	// Встановлюємо таймаут для запиту
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Виконуємо запит GetUser
	res, err := client.GetUser(ctx, &user.UserRequest{Id: "123"})
	if err != nil {
		log.Fatalf("Не вдалося отримати користувача: %v", err)
	}

	fmt.Printf("Отримано користувача: ID=%s, Name=%s\n", res.Id, res.Name)
}
