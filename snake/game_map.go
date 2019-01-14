/*
 * @Author: fylr
 * @Date: 2019-01-13 00:15:51
 * @LastEditors: fylr
 * @LastEditTime: 2019-01-15 00:15:14
 * @Description:
 */

package snake

import (
	"math/rand"
	"time"

	termbox "github.com/nsf/termbox-go"
)

type gameMap struct {
	snake                 *snake
	food                  point
	up, right, down, left int
}

func newGameMap(s *snake, f point, u, r, d, l int) *gameMap {
	return &gameMap{
		snake: s,
		food:  f,
		up:    u,
		right: r,
		down:  d,
		left:  l,
	}
}

func (g *gameMap) flush() int {
	head := g.snake.head()
	newhead := point{head.x + dirToStep[g.snake.direction].x, head.y + dirToStep[g.snake.direction].y}

	var isGrowth = false
	var st = 0

	if newhead.x <= g.left || newhead.x >= g.right || newhead.y <= g.up || newhead.y >= g.down {
		return -1
	} else if newhead.x == g.food.x && newhead.y == g.food.y {
		isGrowth = true
		st = g.snake.move(isGrowth)
		f := g.food
		for !isVaild(g.snake.body, f) {
			x := rand.Intn(g.right-g.left-2) + g.left + 1
			y := rand.Intn(g.down-g.up-2) + g.up + 1
			f = point{x, y}
		}
		g.food = f
	} else {
		st = g.snake.move(isGrowth)
	}
	return st
}

func (g *gameMap) draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for x := g.left; x <= g.right; x++ {
		termbox.SetCell(x, g.up, 0x2500, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(x, g.down, 0x2500, termbox.ColorWhite, termbox.ColorBlack)
	}
	for y := g.up; y <= g.down; y++ {
		termbox.SetCell(g.left, y, 0x2500, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(g.right, y, 0x2500, termbox.ColorWhite, termbox.ColorBlack)
	}

	g.flush()
	for _, temp := range g.snake.body {
		termbox.SetCell(temp.x, temp.y, 0x2500, termbox.ColorRed, termbox.ColorYellow)
	}

	termbox.SetCell(g.food.x, g.food.y, '?', termbox.ColorWhite, termbox.ColorBlue)
}

func initGameMap() *gameMap {
	rand.Seed(time.Now().UnixNano())

	var up, right, down, left int = 5, 85, 35, 5
	var l = 2
	x := rand.Intn(right-left-10) + left + 4
	y := rand.Intn(down-up-10) + up + 4
	dir := rand.Intn(4)
	body := []point{
		point{x - dirToStep[dir].x, y - dirToStep[dir].y},
		point{x, y},
	}
	s := newSnake(body, dir, l)

	f := point{x, y}
	for !isVaild(body, f) {
		x = rand.Intn(right-left-2) + left + 1
		y = rand.Intn(down-up-2) + up + 1
		f = point{x, y}
	}

	return newGameMap(s, f, up, right, down, left)
}

func isVaild(b []point, f point) bool {
	for _, temp := range b {
		if temp.x == f.x && temp.y == f.y {
			return false
		}
	}
	return true
}
