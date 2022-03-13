package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("found empty name")
	}
	message := fmt.Sprintf("Hi %s. Welcome!! ", name)
	return message, nil
}
