package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pixambi/hashicorp-visualised.git/internal/systems/nodes"
	"github.com/pixambi/hashicorp-visualised.git/internal/systems/sprite"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

var vault = sprite.LoadImage("../internal/assets/vault-logo.png")

type Engine struct {
	world       donburi.World
	ecs         *ecs.ECS
	nodeManager *nodes.NodeManager
}

func NewEngine() *Engine {
	world := donburi.NewWorld()
	ecsInstance := ecs.NewECS(world)

	engine := &Engine{
		world:       world,
		ecs:         ecsInstance,
		nodeManager: nodes.NewNodeManager(world),
	}

	//Add Systems
	ecsInstance.AddSystem(engine.updateNodes)
	//Add Renderers
	ecsInstance.AddRenderer(ecs.LayerDefault, engine.drawNodes)

	return engine

}

func (e *Engine) Update() error {
	e.ecs.Update()
	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {
	e.ecs.Draw(screen)
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
