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

type TestScene struct {
	BaseScene
	nodeManager *nodes.NodeManager
}

func NewTestScene() *TestScene {
	return &TestScene{}
}

func (s *TestScene) Enter(world donburi.World, ecs *ecs.ECS) {
	s.BaseScene.Enter(world, ecs)
	s.nodeManager = nodes.NewNodeManager(world)

	s.nodeManager.CreateNode(
		"test",
		"test",
		dmath.Vec2{X: 400, Y: 400},
		"internal/assets/logo/consul-logo.png")

	s.ecs.AddSystem(s.updateNodes)
	s.ecs.AddRenderer(LayerDefault, s.drawNodes)
}

func (s *TestScene) Update(world donburi.World) error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ChangeSignal{Target: "main"}
	}
	s.ecs.Update()
	return nil
}
func (s *TestScene) Draw(screen *ebiten.Image, world donburi.World) {
	s.ecs.Draw(screen)
}

func (s *TestScene) updateNodes(ecs *ecs.ECS) {
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

func (s *TestScene) drawNodes(ecs *ecs.ECS, screen *ebiten.Image) {
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
