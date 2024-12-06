package raft

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pixambi/hashicorp-visualised/scenes"
	"github.com/pixambi/hashicorp-visualised/shapes"
)

type LogoScene struct {
	scenes.BaseScene
	leftLogo   *shapes.VaultLogo
	middleLogo *shapes.VaultLogo
	rightLogo  *shapes.VaultLogo
}

func NewLogoScene() *LogoScene {
	return &LogoScene{}
}

func (s *LogoScene) Init() {
	logoSize := float32(100)
	spacing := float32(200) // Space between logos

	// Calculate starting X position to center the group of logos
	startX := float32(rl.GetScreenWidth())/2 - (spacing)
	centerY := float32(rl.GetScreenHeight())/2 - logoSize/2

	// Create three logos and position them
	s.leftLogo = shapes.NewVaultLogo(logoSize)
	s.middleLogo = shapes.NewVaultLogo(logoSize)
	s.rightLogo = shapes.NewVaultLogo(logoSize)

	// Position logos in a row
	s.leftLogo.MoveTo(startX, centerY)
	s.middleLogo.MoveTo(startX+spacing, centerY)
	s.rightLogo.MoveTo(startX+spacing*2, centerY)
}

func (s *LogoScene) Update() {
	// No updates needed for stationary logos
}

func (s *LogoScene) Draw() {
	// Draw all three logos
	s.leftLogo.Draw()
	s.middleLogo.Draw()
	s.rightLogo.Draw()
}

func (s *LogoScene) Unload() {
	// Clean up any resources if needed
}
