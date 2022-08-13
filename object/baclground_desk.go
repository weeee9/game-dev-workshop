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
	//go:embed images/bg_wood.png
	bgDesk []byte
)

type BackgroundDesk struct {
	img    *ebiten.Image
	width  int
	height int
}

func NewBackgroundDesk(img *ebiten.Image) Object {
	width, height := img.Size()

	return &BackgroundDesk{
		img:    img,
		width:  width,
		height: height,
	}
}

func NewDefaultBackgroundDesk() Object {
	img, _, err := image.Decode(bytes.NewReader(bgDesk))
	if err != nil {
		log.Fatal(err)
	}

	bg := ebiten.NewImageFromImage(img)

	return NewBackgroundDesk(bg)
}

func (bg *BackgroundDesk) Draw(screen *ebiten.Image) error {
	screenWidth, screenHeight := screen.Size()

	colNum := (screenWidth / bg.width) + 1

	y := screenHeight - bg.height/2

	for x := 0; x < colNum; x++ {
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(float64(x*bg.width), float64(y))
		screen.DrawImage(bg.img, opt)
	}

	return nil
}

func (bg *BackgroundDesk) Update() error {
	return nil
}
