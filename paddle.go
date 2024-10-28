package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Paddle struct {
	//TODO
	X     int
	Y     int
	Score int
}

func initPaddle(x int) *Paddle {
	return &Paddle{
		X: x,
		Y: 120,
	}
}

func (paddle *Paddle) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(paddle.X), float32(paddle.Y), paddleW, paddleH, color.RGBA{0x99, 0xcc, 0xff, 0xff}, false)
}

func (paddle *Paddle) Update(dy int) {
	if paddle.Y < 0 && dy < 0 {
		return
	}

	if paddle.Y > screenH-paddleH && dy > 0 {
		return
	}

	paddle.Y += 2 * dy
}
