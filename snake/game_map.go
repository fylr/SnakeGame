/*
 * @Author: fylr
 * @Date: 2019-01-13 00:15:51
 * @LastEditors: fylr
 * @LastEditTime: 2019-01-17 01:12:05
 * @Description:
 */

package snake

import (
	"fmt"
	"math/rand"
	"time"

	termbox "github.com/nsf/termbox-go"
)

type gameMap struct {
	snake *snake
	food  point
	// up, right, down, left int
	left, up, width, height int
	title, helpInfo, status string
	score                   int
}

func newGameMap(snakeVal *snake, foodVal point, leftVal, upVal, widthVal, heightVal int) *gameMap {
	return &gameMap{
		snake:    snakeVal,
		food:     foodVal,
		left:     leftVal,
		up:       upVal,
		width:    widthVal,
		height:   heightVal,
		title:    "Snake Game",
		helpInfo: "↑ ↓ ← → or w s a d  control the direction，p pasue，r restart，q or esc quit the game",
		status:   "run",
		score:    snakeVal.length,
	}
}

func initGameMap() *gameMap {
	rand.Seed(time.Now().UnixNano())

	var left, up, width, height int = 12, 6, 35, 35
	var snakeLen = 2
	x := rand.Intn(width-2*snakeLen) + snakeLen
	y := rand.Intn(height-2*snakeLen) + snakeLen
	dir := rand.Intn(4)
	snakeBody := []point{
		point{x - dirToStep[dir].x, y - dirToStep[dir].y},
		point{x, y},
	}
	snakeVal := newSnake(snakeBody, dir, snakeLen)

	foodVal := point{x, y}
	for !isVaild(snakeBody, foodVal) {
		x = rand.Intn(width-1) + 1
		y = rand.Intn(height-1) + 1
		foodVal = point{x, y}
	}

	return newGameMap(snakeVal, foodVal, left, up, width, height)
}

func (g *gameMap) flush() int {
	snakeHead := g.snake.head()
	nextHead := point{snakeHead.x + dirToStep[g.snake.direction].x, snakeHead.y + dirToStep[g.snake.direction].y}

	var isGrowth = false
	var resultVal = 0

	if nextHead.x <= 0 || nextHead.x > g.width || nextHead.y <= 0 || nextHead.y > g.height {
		resultVal = -1
	} else if !isVaild(g.snake.body, nextHead) {
		resultVal = -2
	} else if nextHead.x == g.food.x && nextHead.y == g.food.y {
		isGrowth = true
		g.score++
		resultVal = g.snake.move(isGrowth)
		foodVal := g.food
		for !isVaild(g.snake.body, foodVal) {
			x := rand.Intn(g.width-1) + 1
			y := rand.Intn(g.height-1) + 1
			foodVal = point{x, y}
		}
		g.food = foodVal
	} else {
		resultVal = g.snake.move(isGrowth)
	}
	if resultVal != 0 {
		g.status = g.snake.dieMsg(resultVal)
	}
	return resultVal
}

func (g *gameMap) draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	//设置边界
	for x := 0; x <= g.width+1; x++ {
		termbox.SetCell(g.left+2*x, g.up, ' ', termbox.ColorBlack, termbox.ColorWhite)
		termbox.SetCell(g.left+2*x+1, g.up, ' ', termbox.ColorBlack, termbox.ColorWhite)

		termbox.SetCell(g.left+2*x, g.up+g.height+1, ' ', termbox.ColorBlack, termbox.ColorWhite)
		termbox.SetCell(g.left+2*x+1, g.up+g.height+1, ' ', termbox.ColorBlack, termbox.ColorWhite)
	}
	for y := 0; y <= g.height+1; y++ {
		termbox.SetCell(g.left, g.up+y, ' ', termbox.ColorBlack, termbox.ColorWhite)
		termbox.SetCell(g.left+1, g.up+y, ' ', termbox.ColorBlack, termbox.ColorWhite)

		termbox.SetCell(g.left+2*(g.width+1), g.up+y, ' ', termbox.ColorBlack, termbox.ColorWhite)
		termbox.SetCell(g.left+2*(g.width+1)+1, g.up+y, ' ', termbox.ColorBlack, termbox.ColorWhite)
	}

	//打印title, help, status, score
	drawMsg(g.left+g.width, 1, termbox.ColorRed, termbox.ColorBlack, g.title)
	drawMsg(g.left, 2, termbox.ColorRed, termbox.ColorBlack, g.helpInfo)
	drawMsg(g.left, 3, termbox.ColorMagenta, termbox.ColorBlack, "status: %s", g.status)
	drawMsg(g.left, 4, termbox.ColorMagenta, termbox.ColorBlack, "scores: %d", g.score-2)

	// g.flush()

	//绘制蛇
	for _, temp := range g.snake.body {
		termbox.SetCell(g.left+2*temp.x, temp.y+g.up, ' ', termbox.ColorBlack, termbox.ColorYellow)
		termbox.SetCell(g.left+2*temp.x+1, temp.y+g.up, ' ', termbox.ColorBlack, termbox.ColorYellow)
	}
	termbox.SetCell(g.left+2*g.snake.head().x, g.snake.head().y+g.up, ' ', termbox.ColorBlack, termbox.ColorRed)
	termbox.SetCell(g.left+2*g.snake.head().x+1, g.snake.head().y+g.up, ' ', termbox.ColorBlack, termbox.ColorRed)

	//食物
	termbox.SetCell(g.left+2*g.food.x, g.food.y+g.up, '?', termbox.ColorWhite, termbox.ColorBlue)
	termbox.SetCell(g.left+2*g.food.x+1, g.food.y+g.up, '?', termbox.ColorWhite, termbox.ColorBlue)
}

func drawMsg(x, y int, fg, bg termbox.Attribute, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func isVaild(b []point, f point) bool {
	for _, temp := range b {
		if temp.x == f.x && temp.y == f.y {
			return false
		}
	}
	return true
}
