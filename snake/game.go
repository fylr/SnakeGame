package snake

import (
	"time"

	termbox "github.com/nsf/termbox-go"
)

//Run func
/**
 * @description:
 * @param {type}
 * @return:
 */
func Run() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	gm := initGameMap()

	var keyval = make(chan int)
	go listenToKeyboard(keyval)

	var isEsc = false
	var isPasue = false
	for !isEsc {
		select {
		case val := <-keyval:
			switch val {
			case 0, 1, 2, 3:
				gm.snake.changeDir(val)
			case 4:
				// isPasue = !isPasue
				if isPasue {
					isPasue = false
					gm.status = "run"
				} else {
					isPasue = true
					gm.status = "pasue"
				}
			case 5:
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				isPasue = false
				gm = initGameMap()
			case 6:
				isEsc = true
			}
		default:
		}
		if !isPasue {
			st := gm.flush()
			if st < 0 {
				isPasue = true
			}
		}
		gm.draw()
		termbox.Flush()
		time.Sleep(time.Duration(500) * time.Millisecond)
	}
}
