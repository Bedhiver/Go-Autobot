package main

import (
	"example/greetings"
	"fmt"
	"log"
	"github.com/go-vgo/robotgo"
	// "rsc.io/quote"
)

func main() {
	robotgo.MouseSleep = 100

	// robotgo.ScrollMouse(10, "up")
	// robotgo.ScrollMouse(20, "right")
  
	// robotgo.Scroll(0, -10)
	// robotgo.Scroll(100, 0)
  
	// robotgo.MilliSleep(100)
	// robotgo.ScrollSmooth(-10, 6)
  
	// robotgo.Move(10, 20)
	// robotgo.MoveRelative(0, -10)
	// robotgo.DragSmooth(10, 10)
  
	// robotgo.Click("wheelRight")
	// robotgo.Click("left", true)
	robotgo.MoveSmooth(100, 200, 1.0, 10.0)
  
	// robotgo.Toggle("left")
	// robotgo.Toggle("left", "up")

	log.SetPrefix("greetings : ")
	log.SetFlags(0)

	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
	// fmt.Println("hello world !")
	// fmt.Println(quote.Go())
}
