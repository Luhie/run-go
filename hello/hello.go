package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("hello Println")
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
}
