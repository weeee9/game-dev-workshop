package object

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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

var (
	//go:embed images/water1.png
	bgWater1 []byte
)

type water struct {
	tick int

	img    *ebiten.Image
	width  int
	height int

	dx         int
	dy         int
	xDirection moveDirection
	yDirection moveDirection
}

func NewWater(img *ebiten.Image) *water {
	width, height := img.Size()

	return &water{
		tick:   0,
		img:    img,
		width:  width,
		height: height,

		xDirection: directionUpOrRight,
		yDirection: directionUpOrRight,
	}
}

func (w *water) Update() error {
	if w.tick > 100 {
		w.tick = 0
	}
	w.tick++

	w.updateDirection()
	w.updateDxDy()

	return nil
}

const (
	xStep = 5
	yStep = 2

	xSpeed = 5
	ySpeed = 10
)

type moveDirection int

const (
	directionUpOrRight  moveDirection = 1
	directionDownOrLeft moveDirection = -1
)

func (w *water) Draw(screen *ebiten.Image) error {
	screenWidth, screenHeight := screen.Size()

	colNum := (screenWidth / w.width) + 1

	sy := screenHeight - (w.height)

	for x := -1; x < colNum+2; x++ {
		opt := &ebiten.DrawImageOptions{}

		sx := x * w.width

		opt.GeoM.Translate(float64(sx+w.dx), float64(sy+w.dy))
		screen.DrawImage(w.img, opt)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("tick: %d\n x:%d, dir: %d\n y:%d, dir: %d", w.tick, w.dx, w.xDirection, w.dy, w.yDirection))

	return nil
}

func (w *water) updateDirection() {
	switch {
	case w.dy >= 20 && w.yDirection == directionUpOrRight:
		w.yDirection = directionDownOrLeft
	case w.dy <= -20 && w.yDirection == directionDownOrLeft:
		w.yDirection = directionUpOrRight

	case w.dx >= 50 && w.xDirection == directionUpOrRight:
		w.xDirection = directionDownOrLeft
	case w.dx <= -50 && w.xDirection == directionDownOrLeft:
		w.xDirection = directionUpOrRight
	}
}

func (w *water) updateDxDy() {
	if w.tick%ySpeed == 0 {
		w.dy += (yStep * int(w.yDirection))
	}

	if w.tick%xSpeed == 0 {
		w.dx += (xStep * int(w.xDirection))
	}
}
