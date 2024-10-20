package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ball struct {
	X    int
	Y    int
	dX   int
	dY   int
	dMax int
}

func initBall() *Ball {
	dMax := 3
	return &Ball{
		X:    160 + (-20 + rand.Intn(40)),
		Y:    rand.Intn(240),
		dX:   dMax,
		dY:   dMax,
		dMax: dMax,
	}

}

func (ball *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(ball.X), float32(ball.Y), 16, 16, color.RGBA{0x99, 0xcc, 0xff, 0xff}, false)
}

func (ball *Ball) Update() {
	ball.X += ball.dX
	ball.Y += ball.dY

	if ball.Y == 0 {
		ball.dY = -ball.dY

	}

	if ball.Y == 240 {
		ball.dY = -ball.dY
	}
}
