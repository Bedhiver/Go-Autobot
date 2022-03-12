package main

import (
	"example/greetings"
	"fmt"
	"rsc.io/quote"
)

func main() {
	message := greetings.Hello("paqui")
	fmt.Println(message)
	fmt.Println("hello world !")
	fmt.Println(quote.Go())
}
