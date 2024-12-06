package entities

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type VoteEntity struct {
	BaseEntity
	size          float32
	speed         float32
	target        rl.Vector2
	reachedTarget bool
}

func NewVoteEntity(col rl.Color, size float32) *VoteEntity {
	return &VoteEntity{
		BaseEntity: BaseEntity{
			position: rl.Vector2{X: 0, Y: 0},
			visible:  true,
		},
		size:  size,
		speed: 5.0, // Adjust this value to change vote travel speed
	}
}

func (v *VoteEntity) SetTarget(target rl.Vector2) {
	v.target = target
}

func (v *VoteEntity) HasReachedTarget() bool {
	return v.reachedTarget
}

func (v *VoteEntity) Update() {
	if v.reachedTarget {
		return
	}

	// Calculate direction vector
	dx := v.target.X - v.position.X
	dy := v.target.Y - v.position.Y

	// Calculate distance to target
	distance := float32(math.Sqrt(float64(dx*dx + dy*dy)))

	if distance < v.speed {
		// Vote has reached its target
		v.position = v.target
		v.reachedTarget = true
		return
	}

	// Normalize direction and apply speed
	v.position.X += (dx / distance) * v.speed
	v.position.Y += (dy / distance) * v.speed
}

func (v *VoteEntity) Draw() {
	if !v.visible {
		return
	}
	rl.DrawCircle(int32(v.position.X), int32(v.position.Y), v.size, rl.Gold)
}
