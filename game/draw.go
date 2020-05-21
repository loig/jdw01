// Package game implements ebiten.Game
/*
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
package game

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

// Draw implements one of the required methods
// for the ebiten.Game interface
func (g *Game) Draw(screen *ebiten.Image) {

	// Draw the blue guy
	op := &ebiten.DrawImageOptions{}
	if g.blueCharacter.facing == left {
		var mirorM ebiten.GeoM
		mirorM.SetElement(0, 0, -1)
		mirorM.SetElement(1, 1, 1)
		op.GeoM.Concat(mirorM)
		op.GeoM.Translate(32, 0)
	}
	op.GeoM.Translate(g.blueCharacter.x, g.blueCharacter.y)
	sub := image.Rect(
		0+32*g.blueCharacter.animationStep, 0+32*int(g.blueCharacter.state),
		32+32*g.blueCharacter.animationStep, 32+32*int(g.blueCharacter.state),
	)
	screen.DrawImage(blueCharacterImage.SubImage(sub).(*ebiten.Image), op)

}
