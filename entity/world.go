package entity

import "github.com/hajimehoshi/ebiten/v2"

type World struct {
	entities []Entity
}

func NewWorld() *World {
	return &World{
		entities: make([]Entity, 0),
	}
}

func (w *World) AddEntity(e Entity) {
	w.entities = append(w.entities, e)
}

func (w *World) Update() error {
	for _, e := range w.entities {
		if err := e.Update(); err != nil {
			return err
		}
	}
	return nil
}

func (w *World) Draw(screen *ebiten.Image) {
	for _, e := range w.entities {
		e.Draw(screen)
	}
}
