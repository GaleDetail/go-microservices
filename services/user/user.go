package main

import (
	"errors"
	"time"
)

type User struct {
	ID   string
	Name string
}

func FetchUserData(userID string) (*User, error) {
	userChan := make(chan *User)
	errorChan := make(chan error)

	go func() {
		time.Sleep(100 * time.Millisecond)
		if userID == "" {
			errorChan <- errors.New("user ID is empty")
			return
		}
		userChan <- &User{ID: userID, Name: "John Doe"}
	}()

	select {
	case user := <-userChan:
		return user, nil
	case err := <-errorChan:
		return nil, err
	}
}
