package sprite

import (
	"github.com/ebitengine/gomobile/asset"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

func LoadImage(path string) *ebiten.Image {
	f, err := asset.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}
