package snake

import "testing"

func TestMoveNoGrowth(t *testing.T) {
	b := []point{point{9, 10}, point{10, 10}}
	s := newSnake(b, 1, 2)

	oldb := []point{point{9, 10}, point{10, 10}}
	for i := range oldb {
		oldb[i].x += dirToStep[1].x
		oldb[i].y += dirToStep[1].y
	}

	st := s.move(false)
	for i, temp := range oldb {
		if temp != s.body[i] {
			st = -3
		}
	}
	if st != 0 || s.length != 2 {
		t.Fatalf("fatal move nogrowth")
	}
}
