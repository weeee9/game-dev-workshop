package object

import "github.com/hajimehoshi/ebiten/v2"

type Object interface {
	Draw(screen *ebiten.Image) error
	Update() error
	OnScreen() bool
}
