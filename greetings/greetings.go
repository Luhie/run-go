package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// 1.          2.      3.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	/*
		var message string
		message = fmt.Sprintf("Hi, %v. Welcom!", name)
		nil =  No error
	*/
	return message, nil
}

// 1.Function name, 2.Parameter type, 3.Retrun type

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcom!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
