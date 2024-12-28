package scene

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pixambi/hashicorp-visualised.git/internal/components"
	"github.com/pixambi/hashicorp-visualised.git/internal/systems/nodes"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

type MainScene struct {
	BaseScene
	nodeManager *nodes.NodeManager
}

func NewMainScene() *MainScene {
	return &MainScene{}
}

func (s *MainScene) Enter(world donburi.World, ecs *ecs.ECS) {
	s.BaseScene.Enter(world, ecs)
	s.nodeManager = nodes.NewNodeManager(world)

	s.nodeManager.CreateNode(
		"test",
		"test",
		dmath.Vec2{X: 400, Y: 400},
		"internal/assets/logo/vault-logo.png")

	s.ecs.AddSystem(s.updateNodes)
	s.ecs.AddRenderer(LayerDefault, s.drawNodes)
}

func (s *MainScene) Update(world donburi.World) error {
	if ebiten.IsKeyPressed(ebiten.Key1) {
		return ChangeSignal{Target: "test"}
	}
	if ebiten.IsKeyPressed(ebiten.Key2) {
		return ChangeSignal{Target: "scene2"}
	}
	if ebiten.IsKeyPressed(ebiten.Key3) {
		return ChangeSignal{Target: "scene3"}
	}

	s.ecs.Update()
	return nil
}
func (s *MainScene) Draw(screen *ebiten.Image, world donburi.World) {
	s.ecs.Draw(screen)
}

func (s *MainScene) updateNodes(ecs *ecs.ECS) {
	query := donburi.NewQuery(filter.Contains(
		components.Node,
		components.Sprite,
		transform.Transform))

	for entry := range query.Iter(s.world) {
		nodeData := components.Node.Get(entry)
		worldPos := transform.WorldPosition(entry)

		switch nodeData.State {
		case "follower":
			fmt.Println(worldPos)
		}
	}
}

func (s *MainScene) drawNodes(ecs *ecs.ECS, screen *ebiten.Image) {
	query := donburi.NewQuery(filter.Contains(
		components.Node,
		transform.Transform,
		components.Sprite,
	))

	for entry := range query.Iter(s.world) {
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
