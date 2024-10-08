package main

import (
	"context"
	"github.com/GaleDetail/go-microservices/api/order"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	order.UnimplementedOrderServiceServer
}

func (s *server) GetOrder(ctx context.Context, req *order.OrderRequest) (*order.OrderResponse, error) {
	order, err := FetchOrderData(req.Id)
	if err != nil {
		return nil, err
	}
	return &order.OrderResponse{Id: order.ID, Description: order.Description}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	order.RegisterOrderServiceServer(s, &server{})
	log.Println("Order service is running on port 50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
