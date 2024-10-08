package main

import "testing"

func TestFetchOrderData(t *testing.T) {
	t.Run("Valid Order ID", func(t *testing.T) {
		order, err := FetchOrderData("456")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if order.ID != "456" {
			t.Errorf("Expected order ID '456', got '%s'", order.ID)
		}
	})

	t.Run("Empty Order ID", func(t *testing.T) {
		_, err := FetchOrderData("")
		if err == nil {
			t.Fatalf("Expected an error, got nil")
		}
	})
}
