package object

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed HUD/crosshair_white_large.png
	crosshairImg []byte

	//go:embed HUD/crosshair_red_large.png
	crosshairClickedImg []byte
)

type Crosshair struct {
	img    *ebiten.Image
	width  int
	height int

	clickedImg       *ebiten.Image
	clickedImgWidth  int
	clickedImgHeight int

	x           int
	y           int
	clicked     bool
	lastClickAt time.Time
}

func NewCrosshair(img, clickedImg *ebiten.Image) Object {
	w, h := img.Size()

	clickedImgW, clickedImgH := clickedImg.Size()

	return &Crosshair{
		img:    img,
		width:  w,
		height: h,

		clickedImg:       clickedImg,
		clickedImgWidth:  clickedImgW,
		clickedImgHeight: clickedImgH,
	}
}

func NewDefaultCrosshair() Object {
	img, _, err := image.Decode(bytes.NewReader(crosshairImg))
	if err != nil {
		log.Fatal(err)
	}

	clickedImg, _, err := image.Decode(bytes.NewReader(crosshairClickedImg))
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	crosshair := ebiten.NewImageFromImage(img)

	crosshairClicked := ebiten.NewImageFromImage(clickedImg)

	return NewCrosshair(crosshair, crosshairClicked)
}

func (c *Crosshair) Draw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.x), float64(c.y))

	if c.clicked {
		op.GeoM.Translate(-float64(c.clickedImgWidth)/2, -float64(c.clickedImgHeight)/2)
		screen.DrawImage(c.clickedImg, op)
		return nil
	}

	op.GeoM.Translate(-float64(c.width)/2, -float64(c.height)/2)
	screen.DrawImage(c.img, op)

	return nil
}

const (
	debouncer = 500 * time.Millisecond
)

func (c *Crosshair) Update() error {
	c.x, c.y = ebiten.CursorPosition()

	c.clicked = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	return nil
}

func (c *Crosshair) OnScreen() bool {
	return true
}
