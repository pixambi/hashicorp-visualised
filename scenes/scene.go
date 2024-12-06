package scenes

// Scene defines the interface for all scenes
type Scene interface {
	Init()
	Update()
	Draw()
	Unload()
	IsActive() bool
	SetActive(bool)
}

// BaseScene provides common functionality for scenes
type BaseScene struct {
	active bool
}

func (b *BaseScene) Init()            {}
func (b *BaseScene) Update()          {}
func (b *BaseScene) Draw()            {}
func (b *BaseScene) Unload()          {}
func (b *BaseScene) IsActive() bool   { return b.active }
func (b *BaseScene) SetActive(a bool) { b.active = a }
