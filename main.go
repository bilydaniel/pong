package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Ball *Ball
}

func (g *Game) Update() error {
	g.Ball.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//ebitenutil.DebugPrint(screen, "Hello, World!")
	if g.Ball != nil {
		g.Ball.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("PONG")

	if err := ebiten.RunGame(&Game{Ball: initBall()}); err != nil {
		log.Fatal(err)
	}
}
