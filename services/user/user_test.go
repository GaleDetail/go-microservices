package main

import "testing"

func TestFetchUserData(t *testing.T) {
	t.Run("Valid User ID", func(t *testing.T) {
		user, err := FetchUserData("123")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if user.ID != "123" {
			t.Errorf("Expected user ID '123', got '%s'", user.ID)
		}
	})

	t.Run("Empty User ID", func(t *testing.T) {
		_, err := FetchUserData("")
		if err == nil {
			t.Fatalf("Expected an error, got nil")
		}
	})
}
