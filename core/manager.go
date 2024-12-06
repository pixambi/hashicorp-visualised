package core

import (
	"github.com/pixambi/hashicorp-visualised/scenes"
)

type SceneManager struct {
	scenes      map[string]scenes.Scene
	activeScene string
}

func NewSceneManager() *SceneManager {
	return &SceneManager{
		scenes: make(map[string]scenes.Scene),
	}
}

func (sm *SceneManager) RegisterScene(name string, scene scenes.Scene) {
	sm.scenes[name] = scene
}

func (sm *SceneManager) SwitchTo(name string) {
	if scene, exists := sm.scenes[name]; exists {
		if current, ok := sm.scenes[sm.activeScene]; ok {
			current.Unload()
		}
		scene.Init()
		sm.activeScene = name
	}
}

func (sm *SceneManager) Update() {
	if scene, exists := sm.scenes[sm.activeScene]; exists {
		scene.Update()
	}
}

func (sm *SceneManager) Draw() {
	if scene, exists := sm.scenes[sm.activeScene]; exists {
		scene.Draw()
	}
}
