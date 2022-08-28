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
	//go:embed images/bg_green.png
	bgGrass []byte
)

type BackgroundGrass struct {
	img    *ebiten.Image
	width  int
	height int
}

func NewBackgroundGrass(img *ebiten.Image) Object {
	width, height := img.Size()

	return &BackgroundGrass{
		img:    img,
		width:  width,
		height: height,
	}
}

func NewDefaultBackgroundGrass() Object {
	img, _, err := image.Decode(bytes.NewReader(bgGrass))
	if err != nil {
		log.Fatal(err)
	}

	bg := ebiten.NewImageFromImage(img)

	return NewBackgroundGrass(bg)
}

func (bg *BackgroundGrass) Draw(screen *ebiten.Image) error {
	screenWidth, screenHeight := screen.Size()

	colNum := (screenWidth / bg.width) + 1
	rowNum := (screenHeight / bg.height) + 1

	for y := 0; y < rowNum; y++ {
		for x := 0; x < colNum; x++ {
			opt := &ebiten.DrawImageOptions{}
			opt.GeoM.Translate(float64(x*bg.width), float64(y*bg.height))
			screen.DrawImage(bg.img, opt)
		}
	}

	return nil
}

func (bg *BackgroundGrass) Update() error {
	return nil
}

func (bg *BackgroundGrass) OnScreen() bool {
	return true
}
