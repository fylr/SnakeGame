package main

import "workspace/snake_game/snake"

func main() {
	snake.Run()

	// var keyval = make(chan int)
	// go listenToKeyboard(keyval)

	// var isEsc = false
	// for !isEsc {
	// 	select {
	// 	case val := <-keyval:
	// 		if val == -1 {
	// 			isEsc = true
	// 		}
	// 	default:
	// 	}
	// }

}

// func listenToKeyboard(keyval chan int) {
// 	termbox.SetInputMode(termbox.InputEsc)

// 	for {
// 		switch ev := termbox.PollEvent(); ev.Type {
// 		case termbox.EventKey:
// 			// fmt.Println(ev.Type)
// 			if ev.Key == termbox.KeyEsc {
// 				keyval <- -1
// 			}
// 		case termbox.EventError:
// 			panic(ev.Err)
// 		}
// 	}
// }
