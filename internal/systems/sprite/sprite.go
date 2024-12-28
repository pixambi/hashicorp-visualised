package sprite

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"os"
)

func LoadImage(path string) *ebiten.Image {
	f, err := os.Open(path)
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
