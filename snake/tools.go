package snake

import termbox "github.com/nsf/termbox-go"

func listenToKeyboard(keyval chan int) {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch {
			case ev.Key == termbox.KeyArrowUp, ev.Ch == 'w':
				keyval <- 0
			case ev.Key == termbox.KeyArrowRight, ev.Ch == 'd':
				keyval <- 1
			case ev.Key == termbox.KeyArrowDown, ev.Ch == 's':
				keyval <- 2
			case ev.Key == termbox.KeyArrowLeft, ev.Ch == 'a':
				keyval <- 3
			case ev.Ch == 'p':
				keyval <- 4
			case ev.Ch == 'r':
				keyval <- 5
			case ev.Key == termbox.KeyEsc, ev.Ch == 'q':
				keyval <- 6
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
