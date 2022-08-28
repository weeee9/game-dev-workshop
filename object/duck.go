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

	remove bool
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

	duck.shoot(ebiten.CursorPosition())

	return nil
}

func (duck *Duck) OnScreen() bool {
	return !duck.remove
}

func (duck *Duck) shoot(clickX, clickY int) bool {
	x := int(duck.dx)
	y := int(duck.dy)

	log.Printf("duck x: %d, x2 %d, y: %d, y2: %d | clickX: %d, clickY: %d",
		x, x+duck.width,
		y, y+duck.height,
		clickX, clickY)

	// Approximate the duck to its rectangle, though there're transparent
	// pixels. For better results we can either approximate the duck to other
	// shapes (like a rectangle+circle) or use image.At() to understand
	// if a transparent pixel was hit
	if (clickX >= x && clickX <= x+duck.width) &&
		(clickY >= y && clickY <= y+duck.height) {
		duck.remove = true
		return true
	}

	return false
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
