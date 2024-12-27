package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pixambi/hashicorp-visualised.git/internal/engine"
)

func main() {
	e := &engine.Engine{}

	err := ebiten.RunGame(e)
	if err != nil {
		panic(err)
	}
}
