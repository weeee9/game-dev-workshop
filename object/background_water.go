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
	tick   int
	waters []Object
}

func NewBackgroundWater(imgs ...*ebiten.Image) Object {

	waters := make([]Object, len(imgs))
	for i, img := range imgs {
		waters[i] = NewWater(img)
	}

	return &BackgroundWater{
		waters: waters,
	}
}

func NewDefaultBackgroundWater() Object {
	water1, _, err := image.Decode(bytes.NewReader(bgWater1))
	if err != nil {
		log.Fatal(err)
	}

	water2, _, err := image.Decode(bytes.NewReader(bgWater2))
	if err != nil {
		log.Fatal(err)
	}

	return NewBackgroundWater(
		ebiten.NewImageFromImage(water1),
		ebiten.NewImageFromImage(water2),
	)
}

func (bg *BackgroundWater) Draw(screen *ebiten.Image) error {
	frameCount := len(bg.waters)

	if frameCount == 0 {
		return nil
	}

	i := (bg.tick / 20) % frameCount

	bg.waters[i].Draw(screen)

	return nil
}

func (bg *BackgroundWater) Update() error {
	bg.tick++
	return nil
}

var (
	//go:embed images/water1.png
	bgWater1 []byte

	//go:embed images/water2.png
	bgWater2 []byte
)

type water struct {
	img    *ebiten.Image
	width  int
	height int
}

func NewWater(img *ebiten.Image) *water {
	width, height := img.Size()

	return &water{
		img:    img,
		width:  width,
		height: height,
	}
}

func (w *water) Update() error {
	return nil
}

func (w *water) Draw(screen *ebiten.Image) error {
	screenWidth, screenHeight := screen.Size()

	colNum := (screenWidth / w.width) + 1

	for x := 0; x < colNum; x++ {
		opt := &ebiten.DrawImageOptions{}

		sx := x * w.width
		sy := screenHeight - (w.height)
		opt.GeoM.Translate(float64(sx), float64(sy))
		screen.DrawImage(w.img, opt)
	}
	return nil
}
