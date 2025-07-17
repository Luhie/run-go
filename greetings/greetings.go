package greetings

import (
	"errors"
	"fmt"
)

// 1.          2.      3.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	/*
		var message string
		message = fmt.Sprintf("Hi, %v. Welcom!", name)
		nil =  No error
	*/
	return message, nil
}

// 1.Function name, 2.Parameter type, 3.Retrun type
