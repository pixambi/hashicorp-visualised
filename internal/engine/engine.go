package engine

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pixambi/hashicorp-visualised.git/internal/systems/nodes"
	"github.com/pixambi/hashicorp-visualised.git/internal/systems/scene"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type Engine struct {
	world        donburi.World
	ecs          *ecs.ECS
	sceneManager *scene.Manager
	nodeManager  *nodes.NodeManager
	screenWidth  int
	screenHeight int
}
type Config struct {
	ScreenWidth  int
	ScreenHeight int
	ResizingMode ebiten.WindowResizingModeType
	InitialScene string
}

func NewEngine(config Config) *Engine {
	world := donburi.NewWorld()
	ecsInstance := ecs.NewECS(world)

	engine := &Engine{
		world:        world,
		ecs:          ecsInstance,
		sceneManager: scene.NewManager(world, ecsInstance),
		nodeManager:  nodes.NewNodeManager(world),
		screenWidth:  config.ScreenWidth,
		screenHeight: config.ScreenHeight,
	}

	engine.registerScenes()

	if err := engine.sceneManager.SwitchScene(config.InitialScene); err != nil {
		panic(err)
	}

	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowResizingMode(config.ResizingMode)

	return engine

}

func (e *Engine) registerScenes() {
	e.sceneManager.RegisterScene("main", scene.NewMainScene())
	e.sceneManager.RegisterScene("test", scene.NewTestScene())
}

func (e *Engine) Update() error {
	if err := e.sceneManager.Update(e.world); err != nil {
		var signal scene.ChangeSignal
		switch {
		case errors.As(err, &signal):
			return e.sceneManager.SwitchScene(signal.Target)
		default:
			return err
		}
	}
	return nil
}

func (e *Engine) Draw(screen *ebiten.Image) {
	e.ecs.Draw(screen)
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
