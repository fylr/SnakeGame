package snake

import "testing"

func TestFlush(t *testing.T) {
	gm := initGameMap()

	oldsnake := newSnake(gm.snake.body, gm.snake.direction, gm.snake.length)
	// oldfood :=point{gm.}
	oldsnake.move(false)
	len := gm.snake.length
	st := gm.flush()

	// for i, temp := range oldsnake.body {
	// 	if temp != gm.snake.body[i] {
	// 		st = -3
	// 	}
	// }

	if st != 0 || len != gm.snake.length {
		t.Fatalf("fatal game map flush")
	}
}
