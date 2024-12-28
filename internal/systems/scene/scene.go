package scene

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type ChangeSignal struct {
	Target string
}

func (c ChangeSignal) Error() string {
	return fmt.Sprintf("change scene to: %s", c.Target)
}

type Scene interface {
	Enter(world donburi.World, ecs *ecs.ECS)
	Exit(world donburi.World)
	Update(world donburi.World) error
	Draw(screen *ebiten.Image, world donburi.World)
}
