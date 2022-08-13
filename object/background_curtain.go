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
	//go:embed images/curtain.png
	bgCurtain []byte
)

type BackgroundCurtain struct {
	img    *ebiten.Image
	width  int
	height int
}

func NewBackgroundCurtain(img *ebiten.Image) Object {
	width, height := img.Size()

	return &BackgroundCurtain{
		img:    img,
		width:  width,
		height: height,
	}
}

func NewDefaultBackgroundCurtain() Object {
	img, _, err := image.Decode(bytes.NewReader(bgCurtain))
	if err != nil {
		log.Fatal(err)
	}

	bg := ebiten.NewImageFromImage(img)

	return NewBackgroundCurtain(bg)
}

func (bg *BackgroundCurtain) Draw(screen *ebiten.Image) error {
	screenWidth, _ := screen.Size()

	leftCurtainOpt := &ebiten.DrawImageOptions{}
	screen.DrawImage(bg.img, leftCurtainOpt)

	rightCurtainOpt := &ebiten.DrawImageOptions{}
	rightCurtainOpt.GeoM.Scale(-1, 1) // flip the curtain image
	rightCurtainOpt.GeoM.Translate(float64(screenWidth), 0)

	screen.DrawImage(bg.img, rightCurtainOpt)

	return nil
}

func (bg *BackgroundCurtain) Update() error {
	return nil
}
