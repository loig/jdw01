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

import "math"

func (g *Game) tryWhiteSpecialMoveUp() {
	switch g.getLaderFieldMove(g.whiteCharacter.x, g.whiteCharacter.y, -g.whiteCharacter.speed/2) {
	case normalFieldMove:
		g.whiteCharacter.y -= g.whiteCharacter.speed * 0.8
		if g.state != tuto2 {
			g.state = whiteSpecialMove
		}
	case endOfLaderFieldMove:
		g.whiteCharacter.y -= g.whiteCharacter.speed * 0.8
		g.whiteCharacter.y = math.Round(g.whiteCharacter.y)
		if g.state != tuto2 {
			g.state = playingWhite
		} else {
			g.whiteCharacter.state = idle
		}
	case noFieldMove:
		g.state = whiteSpecialMoveIdle
	}
}

func (g *Game) tryWhiteSpecialMoveDown() {
	switch g.getLaderFieldMove(g.whiteCharacter.x, g.whiteCharacter.y, +g.whiteCharacter.speed/2) {
	case normalFieldMove:
		g.whiteCharacter.y += g.whiteCharacter.speed * 0.8
		if g.state != tuto2 {
			g.state = whiteSpecialMove
		}
	case endOfLaderFieldMove:
		g.whiteCharacter.y += g.whiteCharacter.speed * 0.8
		g.whiteCharacter.y = math.Round(g.whiteCharacter.y)
		if g.state != tuto2 {
			g.state = playingWhite
		} else {
			g.whiteCharacter.state = idle
		}
	case noFieldMove:
		g.state = whiteSpecialMoveIdle
	}
}

func (g *Game) tryWhiteSpecialSwitchRight() {
	if g.state == whiteSpecialMoveIdle {
		g.state = playingBlue
	}
}

func (g *Game) tryWhiteSpecialSwitchLeft() {
	if g.state == whiteSpecialMoveIdle {
		g.state = playingPink
	}
}
