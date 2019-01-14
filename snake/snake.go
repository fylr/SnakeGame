/*
 * @Author: fylr
 * @Date: 2019-01-12 17:01:18
 * @LastEditors: fylr
 * @LastEditTime: 2019-01-15 00:14:14
 * @Description:
 */

package snake

type point struct {
	x, y int
}

type snake struct {
	body      []point
	direction int
	length    int
}

var dirToStep = [4]point{point{0, -1}, point{1, 0}, point{0, 1}, point{-1, 0}}

func (s *snake) head() point {
	return s.body[s.length-1]
}

func (s *snake) changeDir(dir int) {
	s.direction = dir
}

func (s *snake) move(isGrowth bool) int {
	head := s.head()
	newhead := point{head.x + dirToStep[s.direction].x, head.y + dirToStep[s.direction].y}
	if isGrowth {
		s.body = append(s.body, newhead)
		s.length++
	} else {
		s.body = append(s.body[1:], newhead)
	}

	for _, temp := range s.body[:s.length-1] {
		if newhead.x == temp.x && newhead.y == temp.y {
			return -2
		}
	}
	return 0
}

func newSnake(b []point, d, l int) *snake {
	return &snake{
		body:      b,
		direction: d,
		length:    l,
	}
}

func (s *snake) die(status int) string {
	if status == -1 {
		return "move out game map!"
	} else if status == -2 {
		return "bump into body!"
	}
	return ""
}
