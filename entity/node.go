package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	Update() error
	Draw(screen *ebiten.Image)
	GetPosition() (float64, float64)
	SetPosition(x float64, y float64)
	SetScale(scale float64)
}

type BaseEntity struct {
	x, y    float64
	image   *ebiten.Image
	width   int
	height  int
	scale   float64
	visible bool
	angle   float64
	op      ebiten.DrawImageOptions
}

func NewBaseEntity(img *ebiten.Image, width, height int) *BaseEntity {
	scaledImg := ebiten.NewImage(width, height)

	op := &ebiten.DrawImageOptions{}
	origWidth, origHeight := img.Bounds().Dx(), img.Bounds().Dy()
	scaleX := float64(width) / float64(origWidth)
	scaleY := float64(height) / float64(origHeight)

	op.GeoM.Scale(scaleX, scaleY)
	scaledImg.DrawImage(img, op)

	return &BaseEntity{
		image:   scaledImg,
		width:   width,
		height:  height,
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
	e.op = ebiten.DrawImageOptions{}
	e.op.GeoM.Scale(e.scale, e.scale)

	if e.angle != 0 {
		e.op.GeoM.Rotate(e.angle)
	}

	e.op.GeoM.Translate(-float64(e.width)/2, -float64(e.height)/2)
	e.op.GeoM.Translate(e.x, e.y)

	screen.DrawImage(e.image, &e.op)
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

func (e *BaseEntity) SetRotation(angle float64) {
	e.angle = angle
}

func (e *BaseEntity) SetVisible(visible bool) {
	e.visible = visible
}

func (e *BaseEntity) GetSize() (width, height int) {
	return e.width, e.height
}

func (e *BaseEntity) GetImage() *ebiten.Image {
	return e.image
}
