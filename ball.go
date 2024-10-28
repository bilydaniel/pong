package main

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ball struct {
	X                    float32
	Y                    float32
	dX                   float32
	dY                   float32
	dMax                 float32
	diagonalVectorLength float32
}

func initBall() *Ball {
	return &Ball{
		X:                    float32(160 + (-20 + rand.Intn(40))),
		Y:                    float32(rand.Intn(240)),
		dX:                   ballDMax,
		dY:                   ballDMax,
		dMax:                 ballDMax,
		diagonalVectorLength: float32(math.Hypot(1, 1)),
	}
}

func (ball *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(ball.X), float32(ball.Y), ballH, ballW, color.RGBA{0x99, 0xcc, 0xff, 0xff}, false)
}

func (ball *Ball) Update(player1, player2 *Paddle) {
	if ball.dX >= 2 && ball.dY >= 2 {
		ball.dX /= float32(ball.diagonalVectorLength)
		ball.dY /= float32(ball.diagonalVectorLength)
	}
	ball.X += ball.dX
	ball.Y += ball.dY

	if ball.Y <= 0 {
		ball.dY = -ball.dY

	}

	if ball.Y >= screenH-ballH {
		ball.dY = -ball.dY
	}

	if player1 != nil && player2 != nil {
		//TODO optimize (probably check for distance first, check collision if close)
		if ball.Collides(player1) {
			ball.dX = -ball.dX

		} else if ball.Collides(player2) {
			ball.dX = -ball.dX
		}

	}

	//SCORE
	if ball.X < 0 {
		player2.Score += 1
		ball.Reset()

	}

	if ball.X > screenW {
		player1.Score += 1
		ball.Reset()
	}

}

func (ball *Ball) Collides(player *Paddle) bool {
	return ball.X < float32(player.X+paddleW) &&
		ball.X+ballW > float32(player.X) &&
		ball.Y < float32(player.Y+paddleH) &&
		ball.Y+ballH > float32(player.Y)
}
func (ball *Ball) Reset() {
	*ball = Ball{
		X:                    float32(160 + (-20 + rand.Intn(40))),
		Y:                    float32(rand.Intn(240)),
		dX:                   ballDMax,
		dY:                   ballDMax,
		dMax:                 ballDMax,
		diagonalVectorLength: float32(math.Hypot(1, 1)),
	}
}
