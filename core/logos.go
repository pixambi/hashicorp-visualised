package core

import rl "github.com/gen2brain/raylib-go/raylib"

type DoubleSquare struct {
	x     float32
	y     float32
	size  float32
	color rl.Color
}

func NewDoubleSquare(size float32, color rl.Color) *DoubleSquare {
	return &DoubleSquare{
		x:     0,
		y:     0,
		size:  size,
		color: color,
	}
}

func (d *DoubleSquare) MoveTo(x, y float32) {
	d.x = x
	d.y = y
}

func (d *DoubleSquare) Draw() {
	// Outer square
	rl.DrawRectangle(
		int32(d.x),
		int32(d.y),
		int32(d.size),
		int32(d.size),
		d.color,
	)

	// Inner square (20% smaller than outer)
	padding := d.size * 0.2
	innerSize := d.size - (padding * 2)

	rl.DrawRectangle(
		int32(d.x+padding),
		int32(d.y+padding),
		int32(innerSize),
		int32(innerSize),
		rl.Color{R: 255, G: 255, B: 255, A: 255}, // White color
	)
}
