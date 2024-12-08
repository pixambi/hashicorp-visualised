package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pixambi/hashicorp-visualised/config"
	"github.com/pixambi/hashicorp-visualised/engine"
)

func main() {
	config.Init(1080, 720)

	ebiten.SetWindowSize(config.Current.Width, config.Current.Height)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("HashiCorp Visualised")

	// Create a new instance of engine.Game
	game := engine.NewGame()
	game.PreloadImages()

	// Calculate center point
	centerX := float64(config.Current.Width) / 2
	centerY := float64(config.Current.Height) / 2

	// Triangle configuration
	radius := 150.0 // Distance from center to each point

	// Create top vertex
	topX := centerX
	topY := centerY - radius
	top := game.CreateEntity("vault.png", topX, topY, 60, 60)
	game.AddEntity(top)

	// Create bottom-left vertex
	leftX := centerX - radius*math.Cos(math.Pi/6)
	leftY := centerY + radius*math.Sin(math.Pi/6)
	left := game.CreateEntity("vault.png", leftX, leftY, 60, 60)
	game.AddEntity(left)

	// Create bottom-right vertex
	rightX := centerX + radius*math.Cos(math.Pi/6)
	rightY := centerY + radius*math.Sin(math.Pi/6)
	right := game.CreateEntity("vault.png", rightX, rightY, 60, 60)
	game.AddEntity(right)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
