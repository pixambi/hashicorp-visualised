// entities/vault.go
package entities

import rl "github.com/gen2brain/raylib-go/raylib"

type VaultEntity struct {
	BaseEntity
	size float32
}

func NewVaultEntity(size float32) *VaultEntity {
	return &VaultEntity{
		BaseEntity: BaseEntity{
			position: rl.Vector2{X: 0, Y: 0},
			visible:  true,
		},
		size: size,
	}
}

func (v *VaultEntity) Draw() {
	if !v.visible {
		return
	}

	// Calculate offset to center the triangle around its position point
	halfSize := v.size / 2

	// Draw the triangle centered on position
	rl.DrawTriangle(
		// Top right vertex
		rl.Vector2{X: v.position.X + halfSize, Y: v.position.Y - halfSize},
		// Top left vertex
		rl.Vector2{X: v.position.X - halfSize, Y: v.position.Y - halfSize},
		// Bottom vertex
		rl.Vector2{X: v.position.X, Y: v.position.Y + halfSize},
		rl.Yellow,
	)

	// Draw outline
	rl.DrawTriangleLines(
		rl.Vector2{X: v.position.X - halfSize, Y: v.position.Y - halfSize},
		rl.Vector2{X: v.position.X + halfSize, Y: v.position.Y - halfSize},
		rl.Vector2{X: v.position.X, Y: v.position.Y + halfSize},
		rl.Black,
	)

	// Draw the label
	v.DrawLabel()
}

func (v *VaultEntity) Update() {
	// Add any update logic here
	// For example, animations, state changes, etc.
}
