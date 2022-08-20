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

var (
	//go:embed targets/duck_outline_target_white.png
	duckWhite []byte
)

type Duck struct {
	img    *ebiten.Image
	width  int
	height int

	tick int

	dx         int
	dy         int
	xDirection moveDirection
	yDirection moveDirection
}

func NewDuck(img *ebiten.Image) Object {
	width, height := img.Size()

	return &Duck{
		img:    img,
		width:  width,
		height: height,

		dx: 0 - width,

		xDirection: directionUpOrRight,
		yDirection: directionUpOrRight,
	}
}

func NewDefaultDuck() Object {
	img, _, err := image.Decode(bytes.NewReader(duckWhite))
	if err != nil {
		log.Fatal(err)
	}

	duck := ebiten.NewImageFromImage(img)

	return NewDuck(duck)
}

func (duck *Duck) Draw(screen *ebiten.Image) error {
	screenWidth, screenHeight := screen.Size()

	y := screenHeight - (3 * duck.height)

	if duck.dx > screenWidth {
		duck.dx = 0 - duck.width
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("screen width: %d, duck width: %d\ndx: %d",
		screenWidth, duck.width, duck.dx))

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64(duck.dx), float64(y+duck.dy))

	screen.DrawImage(duck.img, opt)

	return nil
}

func (duck *Duck) Update() error {
	if duck.tick > 100 {
		duck.tick = 0
	}
	duck.tick++

	duck.updateDirection()
	duck.updateDxDy()

	return nil
}

func (duck *Duck) updateDirection() {
	switch {
	case duck.dy >= 10 && duck.yDirection == directionUpOrRight:
		duck.yDirection = directionDownOrLeft
	case duck.dy <= -10 && duck.yDirection == directionDownOrLeft:
		duck.yDirection = directionUpOrRight
	}
}

func (duck *Duck) updateDxDy() {
	if duck.tick%5 == 0 {
		duck.dy += (yStep * int(duck.yDirection))
	}

	if duck.tick%5 == 0 {
		duck.dx += (xStep * int(duck.xDirection))
	}

}
