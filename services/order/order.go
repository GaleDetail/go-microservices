package main

import (
	"errors"
	"time"
)

type Order struct {
	ID          string
	Description string
}

func FetchOrderData(orderID string) (*Order, error) {
	orderChan := make(chan *Order)
	errorChan := make(chan error)

	go func() {
		time.Sleep(100 * time.Millisecond)
		if orderID == "" {
			errorChan <- errors.New("order ID is empty")
			return
		}
		orderChan <- &Order{ID: orderID, Description: "Sample Order"}
	}()

	select {
	case order := <-orderChan:
		return order, nil
	case err := <-errorChan:
		return nil, err
	}
}
