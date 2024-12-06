package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pixambi/hashicorp-visualised/core"
	"github.com/pixambi/hashicorp-visualised/scenes/raft"
)

const (
	screenWidth  = 1280
	screenHeight = 720
	fps          = 60
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Simple Circle")
	defer rl.CloseWindow()

	rl.SetTargetFPS(fps)

	manager := core.NewSceneManager()
	manager.RegisterScene("circle", raft.NewRaftScene())
	manager.SwitchTo("circle")

	for !rl.WindowShouldClose() {
		manager.Update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		manager.Draw()
		rl.EndDrawing()
	}
}
