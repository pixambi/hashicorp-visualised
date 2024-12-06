package raft

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pixambi/hashicorp-visualised/core"
	"github.com/pixambi/hashicorp-visualised/scenes"
)

type LogoScene struct {
	scenes.BaseScene
	logo       *core.DoubleSquare
	movingLogo *core.DoubleSquare
	time       float32
}

func NewLogoScene() *LogoScene {
	return &LogoScene{}
}

func (s *LogoScene) Init() {
	// Create a static logo
	s.logo = core.NewDoubleSquare(100, rl.Red)
	s.logo.MoveTo(100, 100)

	// Create a moving logo
	s.movingLogo = core.NewDoubleSquare(50, rl.Orange)
	s.time = 0
}

func (s *LogoScene) Update() {
	s.time += rl.GetFrameTime()

	// Make the moving logo follow a circular path
	x := float32(400 + 100*math.Cos(float64(s.time)))
	y := float32(300 + 100*math.Sin(float64(s.time)))
	s.movingLogo.MoveTo(x, y)
}

func (s *LogoScene) Draw() {
	// Draw both logos
	s.logo.Draw()
	s.movingLogo.Draw()
}

func (s *LogoScene) Unload() {
	// Clean up any resources if needed
}
