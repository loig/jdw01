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

import "github.com/hajimehoshi/ebiten"

type cameraInfo struct {
	x float64
	y float64
}

func (g *Game) setCameraPosition() {

	switch g.state {
	case playingBlue, blueSpecialMovePhase1, blueSpecialMovePhase2:
		g.camera.x = g.blueCharacter.x * 32
		g.camera.y = g.blueCharacter.y * 32
	case playingPink:
		g.camera.x = g.pinkCharacter.x * 32
		g.camera.y = g.pinkCharacter.y * 32
	case playingWhite:
		g.camera.x = g.whiteCharacter.x * 32
		g.camera.y = g.whiteCharacter.y * 32
	}

}

func (g *Game) applyCamera(op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(
		-g.camera.x+float64(g.screenWidth)/2-16,
		-g.camera.y+float64(g.screenHeight)/2-16,
	)
}
