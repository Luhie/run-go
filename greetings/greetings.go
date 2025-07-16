package greetings

import (
	"fmt"
)

// 1.          2.      3.
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	/*
		var message string
		message = fmt.Sprintf("Hi, %v. Welcom!", name)
	*/
	return message
}

// 1.Function name, 2.Parameter type, 3.Retrun type
