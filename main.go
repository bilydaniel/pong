package main

import (
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

const (
	paddleW  = 8
	paddleH  = 32
	ballH    = 6
	ballW    = 6
	screenW  = 320
	screenH  = 240
	ballDMax = 2
)

type Game struct {
	Ball    *Ball
	Player1 *Paddle
	Player2 *Paddle
}

func (g *Game) Update() error {

	var d1, d2 int

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		d1 -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		d1 += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		d2 -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		d2 += 1
	}

	if g.Player1 != nil {
		g.Player1.Update(d1)
	}

	if g.Player2 != nil {
		g.Player2.Update(d2)
	}

	g.Ball.Update(g.Player1, g.Player2)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//ebitenutil.DebugPrint(screen, "Hello, World!")

	//middle line
	vector.StrokeLine(screen, screenW/2-1, 0, screenW/2-1, screenH, 1, color.White, true)

	text.Draw(screen, strconv.Itoa(g.Player1.Score), basicfont.Face7x13, screenW/2-10, 10, color.White)
	text.Draw(screen, strconv.Itoa(g.Player2.Score), basicfont.Face7x13, screenW/2+2, 10, color.White)

	if g.Ball != nil {
		g.Ball.Draw(screen)
	}

	if g.Player1 != nil {
		g.Player1.Draw(screen)

	}

	if g.Player2 != nil {
		g.Player2.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(screenW), int(screenH)
}

func main() {
	ebiten.SetWindowSize(int(screenW*4), int(screenH*4))
	ebiten.SetWindowTitle("PONG")

	if err := ebiten.RunGame(&Game{
		Ball:    initBall(),
		Player1: initPaddle(10),
		Player2: initPaddle(screenW - 10 - paddleW),
	}); err != nil {
		log.Fatal(err)
	}
}
