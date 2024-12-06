// entities/entity.go
package entities

import rl "github.com/gen2brain/raylib-go/raylib"

// Entity represents any interactive element in the visualization
type Entity interface {
	Draw()
	Update()
	MoveTo(x, y float32)
	GetPosition() rl.Vector2
	GetLabel() string
	SetLabel(string)
	GetState() interface{}
	SetState(interface{})
}

// BaseEntity provides common functionality for all entities
type BaseEntity struct {
	position rl.Vector2
	label    string
	state    interface{}
	visible  bool
}

func (e *BaseEntity) MoveTo(x, y float32) {
	e.position.X = x
	e.position.Y = y
}

func (e *BaseEntity) GetPosition() rl.Vector2 {
	return e.position
}

func (e *BaseEntity) GetLabel() string {
	return e.label
}

func (e *BaseEntity) SetLabel(label string) {
	e.label = label
}

func (e *BaseEntity) GetState() interface{} {
	return e.state
}

func (e *BaseEntity) SetState(state interface{}) {
	e.state = state
}

// DrawLabel draws the entity's label above it
func (e *BaseEntity) DrawLabel() {
	if e.label != "" {
		fontSize := int32(20)
		text := e.label
		textWidth := rl.MeasureText(text, fontSize)
		textX := int32(e.position.X) - textWidth/2
		textY := int32(e.position.Y) - 30 // Offset above the entity
		rl.DrawText(text, textX, textY, fontSize, rl.Black)
	}
}
