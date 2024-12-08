package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2" // Note: you need v2
	"github.com/pixambi/hashicorp-visualised/config"
	"github.com/pixambi/hashicorp-visualised/engine"
)

func main() {
	config.Init(1080, 720)

	ebiten.SetWindowSize(config.Current.WindowWidth, config.Current.WindowHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("HashiCorp Visualised")

	// Create a new instance of engine.Game
	game := engine.NewGame()

	game.PreloadImages()

	centerX := float64(config.Current.DesignWidth) / 2
	centerY := float64(config.Current.DesignHeight) / 2

	top := game.CreateEntity("vault.png", centerX, centerY)
	game.AddEntity(top)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
