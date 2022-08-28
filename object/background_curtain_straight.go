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
	//go:embed images/curtain_straight.png
	bgCurtainStraight []byte
)

type BackgroundCurtainStraight struct {
	img    *ebiten.Image
	width  int
	height int
}

func NewBackgroundCurtainStraight(img *ebiten.Image) Object {
	width, height := img.Size()

	return &BackgroundCurtainStraight{
		img:    img,
		width:  width,
		height: height,
	}
}

func NewDefaultBackgroundCurtainStraight() Object {
	img, _, err := image.Decode(bytes.NewReader(bgCurtainStraight))
	if err != nil {
		log.Fatal(err)
	}

	bg := ebiten.NewImageFromImage(img)

	return NewBackgroundCurtainStraight(bg)
}

func (bg *BackgroundCurtainStraight) Draw(screen *ebiten.Image) error {
	screenWidth, _ := screen.Size()

	colNum := (screenWidth / bg.width) + 1

	for x := 0; x < colNum; x++ {
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(float64(x*bg.width), 0)
		screen.DrawImage(bg.img, opt)
	}

	return nil
}

func (bg *BackgroundCurtainStraight) Update() error {
	return nil
}

func (bg *BackgroundCurtainStraight) OnScreen() bool {
	return true
}
