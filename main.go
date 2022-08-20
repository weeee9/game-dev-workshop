package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/weeee9/game-dev-workshop/game"
)

func main() {
	g := game.NewGame()

	ebiten.SetWindowTitle(g.Title())

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
