package engine

import (
	"github.com/hajimehoshi/ebiten/v2" // Note: you need v2
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game needs to be capitalized to be exported
type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "HashiCorp Visualised")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

