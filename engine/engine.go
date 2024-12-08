package engine

import (
	"github.com/hajimehoshi/ebiten/v2" // Note: you need v2
	"github.com/pixambi/hashicorp-visualised/config"
	"github.com/pixambi/hashicorp-visualised/entity"
	"github.com/pixambi/hashicorp-visualised/resource"
)

// Game needs to be capitalized to be exported
type Game struct {
	world  *entity.World
	images *resource.Manager
}

func NewGame() *Game {
	game := &Game{
		world:  entity.NewWorld(),
		images: resource.NewManager(),
	}
	return game
}

func (g *Game) PreloadImages() {
	g.images.MustLoadImages(
		"vault.png",
	)
}

func (g *Game) Update() error {
	return g.world.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	config.Current.UpdateWindowSize(outsideWidth, outsideHeight)
	return config.Current.DesignWidth, config.Current.DesignHeight
}

// images
// GetImage safely retrieves a loaded image
func (g *Game) GetImage(name string) *ebiten.Image {
	return g.images.GetImage(name)
}

// LoadImage loads a new image if needed
func (g *Game) LoadImage(name string) *ebiten.Image {
	return g.images.LoadImage(name)
}

// Entities

func (g *Game) AddEntity(e entity.Entity) {
	if g.world == nil {
		g.world = entity.NewWorld()
	}
	g.world.AddEntity(e)
}

func (g *Game) CreateEntity(imageName string, x, y float64) entity.Entity {
	img := g.GetImage(imageName)
	e := entity.NewBaseEntity(img)
	e.SetPosition(x, y)
	return e
}
