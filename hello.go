package main

import (
	"example/greetings"
	"fmt"
	"log"
	// "rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings : ")
	log.SetFlags(0)

	message, err := greetings.Hello("ok")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
	// fmt.Println("hello world !")
	// fmt.Println(quote.Go())
}
