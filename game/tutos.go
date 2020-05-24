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
	"github.com/loig/jdw01/world"
)

const (
	idleFrames     = 100
	blueStartFrame = 100 // 200
	pinkStartFrame = 100 // 400
	tuto1EndFrame  = 500
)

func (g *Game) updateTuto1() {
	if g.tutoFrame < tuto1EndFrame {
		if g.tutoFrame > blueStartFrame {
			if g.blueCharacter.state == specialMove {
				g.performBlueSpecialMove()
			} else {
				g.tryRightMove(&g.blueCharacter)
			}
		}
		if g.tutoFrame > idleFrames {
			g.tryRightMove(&g.whiteCharacter)
		}
		if g.tutoFrame > pinkStartFrame {
			if g.pinkCharacter.state == specialMove {
				if g.tutoStep == 0 {
					g.performPinkSpecialMoveUp()
				} else {
					g.performPinkSpecialMoveDown()
				}
			} else {
				g.tryRightMove(&g.pinkCharacter)
			}
		}
	} else {
		g.tutoFrame = 0
		g.blueCharacter.x = world.BlueXTuto1
		g.blueCharacter.y = world.BlueYTuto1
		g.blueCharacter.state = idle
		g.whiteCharacter.x = world.WhiteXTuto1
		g.whiteCharacter.y = world.WhiteYTuto1
		g.whiteCharacter.state = idle
		g.pinkCharacter.x = world.PinkXTuto1
		g.pinkCharacter.y = world.PinkYTuto1
		g.pinkCharacter.state = idle
	}
	g.tutoFrame++
}

const (
	tuto2EndStep1 = 450
	tuto2EndFrame = 700
)

func (g *Game) updateTuto2() {
	if g.tutoFrame < tuto2EndFrame {
		if g.tutoFrame > idleFrames {
			if g.tutoFrame < tuto2EndStep1 {
				g.tutoStep = 0
				if g.whiteCharacter.state == specialMove {
					g.tryWhiteSpecialMoveUp()
				} else {
					g.tryUpMove(&g.whiteCharacter)
				}
				if g.pinkCharacter.state == specialMove {
					g.performPinkSpecialMoveDirectUp()
				} else {
					g.tryUpMove(&g.pinkCharacter)
				}
				g.tryUpMove(&g.blueCharacter)
			} else {
				g.tutoStep = 1
				if g.whiteCharacter.state == specialMove {
					g.tryWhiteSpecialMoveDown()
				} else {
					g.tryDownMove(&g.whiteCharacter)
				}
				if g.pinkCharacter.state == specialMove {
					g.performPinkSpecialMoveDirectDown()
				} else {
					g.tryDownMove(&g.pinkCharacter)
				}
				g.tryDownMove(&g.blueCharacter)
			}
		}
	} else {
		g.tutoFrame = 0
		g.blueCharacter.x = world.BlueXTuto2
		g.blueCharacter.y = world.BlueYTuto2
		g.blueCharacter.state = idle
		g.whiteCharacter.x = world.WhiteXTuto2
		g.whiteCharacter.y = world.WhiteYTuto2
		g.whiteCharacter.state = idle
		g.pinkCharacter.x = world.PinkXTuto2
		g.pinkCharacter.y = world.PinkYTuto2
		g.pinkCharacter.state = idle
	}
	g.tutoFrame++
}

const (
	tuto3EndFrame = 200
)

func (g *Game) updateTuto3() {
	if g.tutoFrame < tuto3EndFrame {
		if g.tutoFrame > idleFrames {
			if g.tutoStep == 0 {
				g.whiteCharacter.state = strike
				g.whiteCharacter.strikeCurrentFrame = 0
				g.blueCharacter.state = strike
				g.blueCharacter.strikeCurrentFrame = 0
				g.pinkCharacter.state = strike
				g.pinkCharacter.strikeCurrentFrame = 0
				g.tutoStep = 1
			} else if g.tutoStep == 1 {
				g.whiteCharacter.strikeCurrentFrame++
				if g.whiteCharacter.strikeCurrentFrame >= g.blueCharacter.strikeNumFrames {
					g.strikeEffectOnField(g.whiteCharacter.x, g.whiteCharacter.y, 1)
					g.whiteCharacter.state = idle
					g.tutoStep = 2
				}
				g.blueCharacter.strikeCurrentFrame++
				if g.blueCharacter.strikeCurrentFrame >= g.blueCharacter.strikeNumFrames {
					g.strikeEffectOnField(g.blueCharacter.x, g.blueCharacter.y, 1)
					g.blueCharacter.state = idle
					g.tutoStep = 2
				}
				g.pinkCharacter.strikeCurrentFrame++
				if g.pinkCharacter.strikeCurrentFrame >= g.blueCharacter.strikeNumFrames {
					g.strikeEffectOnField(g.pinkCharacter.x, g.pinkCharacter.y, 1)
					g.pinkCharacter.state = idle
					g.tutoStep = 2
				}
			}
		}
	} else {
		g.tutoFrame = 0
		g.tutoStep = 0
		g.field = world.GetTuto3Field()
	}
	g.tutoFrame++
}

func (g *Game) updateTuto4() {
	if g.tutoFrame > idleFrames {
		g.tutoFrame = 0
		g.tutoStep = (g.tutoStep + 1) % 3
	}
	g.tutoFrame++
}
