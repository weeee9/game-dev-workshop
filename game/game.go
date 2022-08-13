package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/weeee9/game-dev-workshop/object"
)

type Game struct {
	objects []object.Object
}

func NewGame() *Game {
	return &Game{
		objects: []object.Object{
			object.NewDefaultBackgroundGrass(),
			object.NewDefaultBackgroundWater(),
			object.NewDefaultBackgroundDesk(),
			object.NewDefaultBackgroundCurtain(),
			object.NewDefaultBackgroundCurtainStraight(),
		},
	}
}

func (g *Game) Update() error {
	for _, obj := range g.objects {
		obj.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, obj := range g.objects {
		obj.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}
