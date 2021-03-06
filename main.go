/*
A game
Copyright (C) 2020  Loïg Jezequel

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/loig/jdw01/game"
)

func main() {

	g := &game.Game{}

	// initialize the game
	loadError := g.Init()
	if loadError != nil {
		panic(loadError)
	}

	// initialize window
	ebiten.SetWindowSize(1024, 512)
	ebiten.SetWindowTitle("BWP")

	// run the game
	if err := ebiten.RunGame(g); err != nil {
		return
	}

}
