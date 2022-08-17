package object

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type BackgroundWater struct {
	water *water
}

func NewBackgroundWater(img *ebiten.Image) Object {
	return &BackgroundWater{
		water: NewWater(img),
	}
}

func NewDefaultBackgroundWater() Object {
	waterImg, _, err := image.Decode(bytes.NewReader(bgWater1))
	if err != nil {
		log.Fatal(err)
	}

	return NewBackgroundWater(
		ebiten.NewImageFromImage(waterImg),
	)
}

func (bg *BackgroundWater) Draw(screen *ebiten.Image) error {
	bg.water.Draw(screen)

	return nil
}

func (bg *BackgroundWater) Update() error {
	bg.water.Update()
	return nil
}

const (
	defaultWaterSpeed = 7
)

var (
	//go:embed images/water1.png
	bgWater1 []byte
)

type water struct {
	tick  int
	speed int

	img    *ebiten.Image
	width  int
	height int

	dy         int
	xDirection int
	yDirection int
}

func NewWater(img *ebiten.Image) *water {
	width, height := img.Size()

	return &water{
		tick:   0,
		speed:  defaultWaterSpeed,
		img:    img,
		width:  width,
		height: height,

		dy:         0,
		xDirection: 1,
		yDirection: 1,
	}
}

func (w *water) Update() error {
	w.tick++
	return nil
}

const (
	step = 2
)

func (w *water) Draw(screen *ebiten.Image) error {
	screenWidth, screenHeight := screen.Size()

	colNum := (screenWidth / w.width) + 1

	sy := screenHeight - (w.height)

	if w.dy == -20 || w.dy == 20 {
		w.yDirection *= -1
	}

	if w.tick%w.speed == 0 {
		w.dy += (step * w.yDirection)
	}

	for x := 0; x < colNum; x++ {
		opt := &ebiten.DrawImageOptions{}

		sx := x * w.width

		opt.GeoM.Translate(float64(sx), float64(sy+w.dy))
		screen.DrawImage(w.img, opt)
	}

	return nil
}
