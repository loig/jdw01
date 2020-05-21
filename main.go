package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/loig/jdw01/game"
)

func main() {

	g := &game.Game{}
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}

}
