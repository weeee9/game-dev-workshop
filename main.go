package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/weeee9/game-dev-workshop/game"
)

func main() {
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
