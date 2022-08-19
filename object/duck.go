package object

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed targets/duck_outline_target_white.png
	duckWhite []byte
)

type Duck struct {
	img    *ebiten.Image
	width  int
	height int
}

func NewDuck(img *ebiten.Image) Object {
	width, height := img.Size()

	return &Duck{
		img:    img,
		width:  width,
		height: height,
	}
}

func NewDefaultDuck() Object {
	img, _, err := image.Decode(bytes.NewReader(duckWhite))
	if err != nil {
		log.Fatal(err)
	}

	bg := ebiten.NewImageFromImage(img)

	return NewDuck(bg)
}

func (bg *Duck) Draw(screen *ebiten.Image) error {
	_, screenHeight := screen.Size()

	y := screenHeight - (3 * bg.height)

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(10, float64(y))

	screen.DrawImage(bg.img, opt)

	return nil
}

func (bg *Duck) Update() error {
	return nil
}
