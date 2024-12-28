package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type BaseScene struct {
	world donburi.World
	ecs   *ecs.ECS
}

func (s *BaseScene) Enter(world donburi.World, ecs *ecs.ECS) {
	s.world = world
	s.ecs = ecs
}
func (s *BaseScene) Exit(world donburi.World) {}

func (s *BaseScene) Update(world donburi.World) error {
	return nil
}

func (s *BaseScene) Draw(screen *ebiten.Image, world donburi.World) {
	s.ecs.Draw(screen)
}

func (s *BaseScene) GetWorld() donburi.World {
	return s.world
}

func (s *BaseScene) GetECS() *ecs.ECS {
	return s.ecs
}
