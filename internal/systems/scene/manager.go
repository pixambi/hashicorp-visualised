package scene

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type Manager struct {
	world         donburi.World
	ecs           *ecs.ECS
	scenes        map[string]Scene
	currentScene  string
	previousScene string
}

func NewManager(world donburi.World, ecs *ecs.ECS) *Manager {
	return &Manager{
		world:  world,
		ecs:    ecs,
		scenes: make(map[string]Scene),
	}
}

func (m *Manager) RegisterScene(name string, scene Scene) {
	m.scenes[name] = scene
}

func (m *Manager) SwitchScene(name string) error {
	if scene, exists := m.scenes[name]; exists {
		if currentScene, ok := m.scenes[m.currentScene]; ok {
			currentScene.Exit(m.world)
		}

		m.previousScene = m.currentScene
		m.currentScene = name
		scene.Enter(m.world, m.ecs)
		return nil
	}
	return fmt.Errorf("scene %s not found", name)
}

func (m *Manager) ReturnToPreviousScene() error {
	if m.previousScene != "" {
		return m.SwitchScene(m.previousScene)
	}
	return fmt.Errorf("no previous scene")
}

func (m *Manager) Update(world donburi.World) error {
	if scene, exists := m.scenes[m.currentScene]; exists {
		return scene.Update(m.world)
	}
	return nil
}

func (m *Manager) Draw(screen *ebiten.Image) {
	if scene, exists := m.scenes[m.currentScene]; exists {
		scene.Draw(screen, m.world)
	}
}
