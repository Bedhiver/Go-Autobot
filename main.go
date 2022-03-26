package main

import (
	// "example/greetings"
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	// "log"
	// "rsc.io/quote"
)

func main() {

	add()

	// robotgo.MouseSleep = 100
	// hook.Register(hook.KeyDown, []string{"e"}, func(e hook.Event) {
	// 	fmt.Println("key pressed was e")
	// 	hook.End()
	// })

	// s := hook.Start()
	// <-hook.Process(s)

	// condition := false

	// for {
	// 	// eventPressed := robotgo.AddEvent("e")
	// 	// if eventPressed {
	// 	// 	break
	// 	// }
	// 	robotgo.MoveSmooth(500, 500, 1.0, 1.0)
	// 	robotgo.MoveSmooth(1000, 500, 1.0, 1.0)
	// 	// fmt.Println(robotgo.GetMousePos())
	// 	if condition {
	// 		break
	// 	}
	// }

	fmt.Println("Press any key to exit")
	// var input string
	// fmt.Scanf(input, "%s")

	// log.SetPrefix("greetings : ")
	// log.SetFlags(0)

	// message, err := greetings.Hello("d")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(message)
	// fmt.Println("hello world !")
	// fmt.Println(quote.Go())
}

func add() {

	// condition := false

	s := hook.Start()
	defer hook.End()

	// ct := false

	for {

		i := <-s

		// fmt.Println(i)
		defer robotgo.MoveSmooth(500, 500, 1.0, 1.0)
		robotgo.MoveSmooth(1000, 500, 1.0, 1.0)

		// if i.Kind == hook.KeyHold && i.Rawcode == 59 {
		if i.Rawcode == 101 {
			fmt.Println("HELLLOOOOOO")
			// ct = true

			break
		}
		// if ct && i.Rawcode == 12 {
		// if ct {
		// 	break
		// }
	}

}
