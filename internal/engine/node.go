package engine

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pixambi/hashicorp-visualised.git/internal/components"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

func (e *Engine) updateNodes(ecs *ecs.ECS) {
	query := donburi.NewQuery(filter.Contains(
		components.Node,
		components.Sprite,
		transform.Transform))

	for entry := range query.Iter(e.world) {
		nodeData := components.Node.Get(entry)
		worldPos := transform.WorldPosition(entry)

		switch nodeData.State {
		case "follower":
			fmt.Println(worldPos)
		}
	}
}

func (e *Engine) drawNodes(screen *ebiten.Image, ecs *ecs.ECS) {
	query := donburi.NewQuery(filter.Contains(
		components.Node,
		transform.Transform,
		components.Sprite,
	))
	for entry := range query.Iter(e.world) {
		spriteData := components.Sprite.Get(entry)
		if spriteData.Image == nil {
			continue
		}

		worldPos := transform.WorldPosition(entry)
		worldRotation := transform.WorldRotation(entry)
		worldScale := transform.WorldScale(entry)

		op := &ebiten.DrawImageOptions{}

		op.GeoM.Translate(-float64(spriteData.Image.Bounds().Dx())/2, -float64(spriteData.Image.Bounds().Dy())/2)
		op.GeoM.Scale(worldScale.X*spriteData.Scale, worldScale.Y*spriteData.Scale)
		op.GeoM.Rotate(worldRotation)
		op.GeoM.Translate(worldPos.X, worldPos.Y)

		screen.DrawImage(spriteData.Image, op)
	}
}
