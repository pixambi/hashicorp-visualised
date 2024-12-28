package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pixambi/hashicorp-visualised.git/internal/engine"
)

func main() {

	config := engine.Config{
		ScreenWidth:  800,
		ScreenHeight: 600,
		ResizingMode: ebiten.WindowResizingModeEnabled,
		InitialScene: "main",
	}

	err := ebiten.RunGame(engine.NewEngine(config))
	if err != nil {
		panic(err)
	}
}
