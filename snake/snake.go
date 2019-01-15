/*
 * @Author: fylr
 * @Date: 2019-01-12 17:01:18
 * @LastEditors: fylr
 * @LastEditTime: 2019-01-15 23:07:59
 * @Description:
 */

package snake

var dirToStep = [4]point{point{0, -1}, point{1, 0}, point{0, 1}, point{-1, 0}}

type point struct {
	x, y int
}

type snake struct {
	body      []point
	direction int
	length    int
}

func newSnake(bodyVal []point, dirVal, lenVal int) *snake {
	return &snake{
		body:      bodyVal,
		direction: dirVal,
		length:    lenVal,
	}
}

func (s *snake) head() point {
	return s.body[s.length-1]
}

func (s *snake) changeDir(dir int) {
	s.direction = dir
}

func (s *snake) move(isGrowth bool) int {
	head := s.head()
	nextHead := point{head.x + dirToStep[s.direction].x, head.y + dirToStep[s.direction].y}
	if isGrowth {
		s.body = append(s.body, nextHead)
		s.length++
	} else {
		s.body = append(s.body[1:], nextHead)
	}

	for _, temp := range s.body[:s.length-1] {
		if nextHead.x == temp.x && nextHead.y == temp.y {
			return -2
		}
	}
	return 0
}

func (s *snake) die(status int) string {
	if status == -1 {
		return "move out game map!"
	} else if status == -2 {
		return "bump into body!"
	}
	return ""
}
