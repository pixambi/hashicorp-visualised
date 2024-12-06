package shapes

import rl "github.com/gen2brain/raylib-go/raylib"

type VaultLogo struct {
	x    float32
	y    float32
	size float32
}

func NewVaultLogo(size float32) *VaultLogo {
	return &VaultLogo{
		x:    0,
		y:    0,
		size: size,
	}
}

func (v *VaultLogo) MoveTo(x, y float32) {
	v.x = x
	v.y = y
}

func (v *VaultLogo) Draw() {
	// Draw filled yellow triangle
	rl.DrawTriangle(
		rl.Vector2{X: v.x + v.size, Y: v.y},            // Right vertex
		rl.Vector2{X: v.x, Y: v.y},                     // Left vertex
		rl.Vector2{X: v.x + v.size/2, Y: v.y + v.size}, // Bottom vertex
		rl.Yellow,
	)

	// Draw black outline
	rl.DrawTriangleLines(
		rl.Vector2{X: v.x, Y: v.y},                     // Left vertex
		rl.Vector2{X: v.x + v.size, Y: v.y},            // Right vertex
		rl.Vector2{X: v.x + v.size/2, Y: v.y + v.size}, // Bottom vertex
		rl.Black,
	)
}
