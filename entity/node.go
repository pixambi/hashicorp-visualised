package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pixambi/hashicorp-visualised/config"
)

type Entity interface {
	Update() error
	Draw(screen *ebiten.Image)
	GetPosition() (float64, float64)
	SetPosition(x float64, y float64)
}

type BaseEntity struct {
	x, y    float64
	image   *ebiten.Image
	scale   float64
	visible bool
}

func NewBaseEntity(img *ebiten.Image) *BaseEntity {
	return &BaseEntity{
		image:   img,
		scale:   1.0,
		visible: true,
	}
}

func (e *BaseEntity) Update() error {
	return nil
}

func (e *BaseEntity) Draw(screen *ebiten.Image) {
	if !e.visible || e.image == nil {
		return
	}

	op := &ebiten.DrawImageOptions{}

	// Get image dimensions
	w, h := float64(e.image.Bounds().Dx()), float64(e.image.Bounds().Dy())

	// Scale the image
	op.GeoM.Scale(e.scale, e.scale)

	// Get scaled screen position
	screenX, screenY := config.Current.ScalePosition(e.x, e.y)

	// Center the image at its position
	op.GeoM.Translate(-w*e.scale/2, -h*e.scale/2)
	op.GeoM.Translate(screenX, screenY)

	screen.DrawImage(e.image, op)
}

func (e *BaseEntity) GetPosition() (float64, float64) {
	return e.x, e.y
}

func (e *BaseEntity) SetPosition(x, y float64) {
	e.x = x
	e.y = y
}

func (e *BaseEntity) SetScale(scale float64) {
	e.scale = scale
}

func (e *BaseEntity) SetVisible(visible bool) {
	e.visible = visible
}
