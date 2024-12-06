package raft

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/pixambi/hashicorp-visualised/entities"
	"github.com/pixambi/hashicorp-visualised/scenes"
)

type RaftScene struct {
	scenes.BaseScene
	nodes           [3]*entities.VaultEntity
	electionStarted bool
}

func NewRaftScene() *RaftScene {
	return &RaftScene{}
}

func (s *RaftScene) Init() {
	logoSize := float32(100)
	radius := float32(200) // Distance from center to each node
	numNodes := 3

	// Calculate center of screen
	centerX := float32(rl.GetScreenWidth()) / 2
	centerY := float32(rl.GetScreenHeight()) / 2

	// Create three Vault entities
	for i := 0; i < numNodes; i++ {
		s.nodes[i] = entities.NewVaultEntity(logoSize, numNodes)

		// Calculate position using trigonometry
		angle := float32(2*math.Pi*float64(i)/3 - math.Pi/2)
		x := centerX + float32(math.Cos(float64(angle)))*radius
		y := centerY + float32(math.Sin(float64(angle)))*radius

		s.nodes[i].MoveTo(x, y)
		s.nodes[i].SetLabel("Candidate")
		s.nodes[i].SetState(entities.Candidate)
	}
}

func (s *RaftScene) Update() {
	// Start the election process if it hasn't started
	if !s.electionStarted {
		s.startElection()
		s.electionStarted = true
	}

	// Update all nodes
	for _, node := range s.nodes {
		node.Update()
	}
}

func (s *RaftScene) startElection() {
	// Each node sends votes to others
	for i, sender := range s.nodes {
		for j, receiver := range s.nodes {
			if i != j { // Don't send vote to self
				sender.SendVoteTo(receiver)
			}
		}
	}
}

func (s *RaftScene) Draw() {
	// Draw connecting lines between nodes
	for i := 0; i < len(s.nodes); i++ {
		for j := i + 1; j < len(s.nodes); j++ {
			pos1 := s.nodes[i].GetPosition()
			pos2 := s.nodes[j].GetPosition()
			rl.DrawLineV(pos1, pos2, rl.Gray)
		}
	}

	// Draw all votes
	for _, node := range s.nodes {
		node.DrawVotes()
	}

	// Draw all nodes (on top of votes and lines)
	for _, node := range s.nodes {
		node.Draw()
	}
}

func (s *RaftScene) Unload() {
	// Clean up any resources if needed
}
