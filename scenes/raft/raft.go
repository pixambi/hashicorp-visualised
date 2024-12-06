package raft

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pixambi/hashicorp-visualised/entities"
	"github.com/pixambi/hashicorp-visualised/scenes"
)

type RaftScene struct {
	scenes.BaseScene
	nodes [3]*entities.VaultEntity
}

func NewRaftScene() *RaftScene {
	return &RaftScene{}
}

func (s *RaftScene) Init() {
	logoSize := float32(100)
	radius := float32(200) // Distance from center to each node

	// Calculate center of screen
	centerX := float32(rl.GetScreenWidth()) / 2
	centerY := float32(rl.GetScreenHeight()) / 2

	// Create three Vault entities
	for i := 0; i < 3; i++ {
		s.nodes[i] = entities.NewVaultEntity(logoSize)

		// Calculate position using trigonometry
		// 2π/3 = 120 degrees, offset by -90 degrees (-π/2) to start at top
		angle := float32(2*math.Pi*float64(i)/3 - math.Pi/2)
		x := centerX + float32(math.Cos(float64(angle)))*radius
		y := centerY + float32(math.Sin(float64(angle)))*radius

		s.nodes[i].MoveTo(x, y)

		// Set initial labels
		switch i {
		case 0:
			s.nodes[i].SetLabel("Follower")
			s.nodes[i].SetState("FOLLOWER")
		case 1:
			s.nodes[i].SetLabel("Follower")
			s.nodes[i].SetState("FOLLOWER")
		case 2:
			s.nodes[i].SetLabel("Follower")
			s.nodes[i].SetState("FOLLOWER")
		}
	}
}

func (s *RaftScene) Update() {
	for _, node := range s.nodes {
		node.Update()
	}
}

func (s *RaftScene) Draw() {
	// Draw connecting lines first (behind the nodes)

	// Draw all nodes
	for _, node := range s.nodes {
		node.Draw()
	}
}

func (s *RaftScene) Unload() {
	// Clean up any resources if needed
}
