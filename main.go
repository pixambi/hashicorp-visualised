package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2" // Note: you need v2
	"github.com/pixambi/hashicorp-visualised/engine"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("HashiCorp Visualised")

	// Create a new instance of engine.Game
	game := &engine.Game{}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
