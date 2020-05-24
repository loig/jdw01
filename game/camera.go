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
	"math"

	"github.com/hajimehoshi/ebiten"
)

type cameraInfo struct {
	x float64
	y float64
}

func (g *Game) setCameraPosition() {

	switch g.state {
	case playingBlue, blueSpecialMove:
		g.camera.x = g.blueCharacter.x * 32
		g.camera.y = g.blueCharacter.y * 32
	case playingPink, pinkSpecialMoveUp, pinkSpecialMoveDown, pinkSpecialMoveDirectUp, pinkSpecialMoveDirectDown:
		g.camera.x = g.pinkCharacter.x * 32
		g.camera.y = g.pinkCharacter.y * 32
	case playingWhite, whiteSpecialMove, whiteSpecialMoveIdle:
		g.camera.x = g.whiteCharacter.x * 32
		g.camera.y = g.whiteCharacter.y * 32
	case tuto1, tuto2, tuto3:
		g.camera.x = (float64(len(g.field[0])) / 2) * 32
		g.camera.y = (float64(len(g.field)) / 2) * 32
	case tuto4:
		switch g.tutoStep {
		case 0:
			g.camera.x = g.blueCharacter.x * 32
			g.camera.y = g.blueCharacter.y * 32
		case 1:
			g.camera.x = g.pinkCharacter.x * 32
			g.camera.y = g.pinkCharacter.y * 32
		case 2:
			g.camera.x = g.whiteCharacter.x * 32
			g.camera.y = g.whiteCharacter.y * 32
		}

	}

}

func (g *Game) applyCamera(op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(
		-g.camera.x+float64(g.screenWidth)/2-16,
		-g.camera.y+float64(g.screenHeight)/2-16,
	)
}

func (g *Game) visibleRectangle() (xmin, ymin, xmax, ymax int) {
	xmin = int(math.Round((g.camera.x-float64(g.screenWidth)/2)/32)) - 1
	ymin = int(math.Round((g.camera.y-float64(g.screenHeight)/2)/32)) - 1
	xmax = int(math.Round((g.camera.x+float64(g.screenWidth)/2)/32)) + 1
	ymax = int(math.Round((g.camera.y+float64(g.screenHeight)/2)/32)) + 1
	if xmin < 0 {
		xmin = 0
	}
	if ymin < 0 {
		ymin = 0
	}
	if ymax > len(g.field) {
		ymax = len(g.field)
	}
	if xmax > len(g.field[0]) {
		xmax = len(g.field[0])
	}
	return xmin, ymin, xmax, ymax
}
