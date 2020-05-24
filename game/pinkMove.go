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

	"github.com/loig/jdw01/util"
)

func (g *Game) performPinkSpecialMoveDirectDown() {
	g.pinkCharacter.specialMoveCurrentFrame++
	// vertical move
	firstSplit := g.pinkCharacter.specialMoveNumFrames / 3
	secondSplit := g.pinkCharacter.specialMoveNumFrames - firstSplit
	if g.pinkCharacter.specialMoveCurrentFrame <= firstSplit {
		t := float64(g.pinkCharacter.specialMoveCurrentFrame) / float64(firstSplit)
		tprev := float64(g.pinkCharacter.specialMoveCurrentFrame-1) / float64(firstSplit)
		frameVerticalMove := util.SmoothStop6(t) - util.SmoothStop6(tprev)
		g.pinkCharacter.y -= 0.25 * frameVerticalMove
	} else {
		t := float64(g.pinkCharacter.specialMoveCurrentFrame-firstSplit) / float64(secondSplit)
		tprev := float64(g.pinkCharacter.specialMoveCurrentFrame-1-firstSplit) / float64(secondSplit)
		frameVerticalMove := util.SmoothStart3(t) - util.SmoothStart3(tprev)
		g.pinkCharacter.y += 1.25 * frameVerticalMove
	}
	// end of move
	if g.pinkCharacter.specialMoveCurrentFrame >= g.pinkCharacter.specialMoveNumFrames {
		g.state = playingPink
		g.pinkCharacter.specialMoveCurrentFrame = 0
		g.pinkCharacter.state = idle
		g.pinkCharacter.y = math.Round(g.pinkCharacter.y)
	}
}

func (g *Game) performPinkSpecialMoveDirectUp() {
	g.pinkCharacter.specialMoveCurrentFrame++
	// vertical move
	firstSplit := g.pinkCharacter.specialMoveNumFrames / 3
	secondSplit := g.pinkCharacter.specialMoveNumFrames - firstSplit
	if g.pinkCharacter.specialMoveCurrentFrame <= firstSplit {
		t := float64(g.pinkCharacter.specialMoveCurrentFrame) / float64(firstSplit)
		tprev := float64(g.pinkCharacter.specialMoveCurrentFrame-1) / float64(firstSplit)
		frameVerticalMove := util.SmoothStop3(t) - util.SmoothStop3(tprev)
		g.pinkCharacter.y -= 1.25 * frameVerticalMove
	} else {
		t := float64(g.pinkCharacter.specialMoveCurrentFrame-firstSplit) / float64(secondSplit)
		tprev := float64(g.pinkCharacter.specialMoveCurrentFrame-1-firstSplit) / float64(secondSplit)
		frameVerticalMove := util.SmoothStart3(t) - util.SmoothStart3(tprev)
		g.pinkCharacter.y += 0.25 * frameVerticalMove
	}
	if g.pinkCharacter.specialMoveCurrentFrame >= g.pinkCharacter.specialMoveNumFrames {
		g.state = playingPink
		g.pinkCharacter.specialMoveCurrentFrame = 0
		g.pinkCharacter.state = idle
		g.pinkCharacter.y = math.Round(g.pinkCharacter.y)
	}
}

func (g *Game) performPinkSpecialMoveDown() {
	g.pinkCharacter.specialMoveCurrentFrame++
	// horizontal move
	t := float64(g.pinkCharacter.specialMoveCurrentFrame) / float64(g.pinkCharacter.specialMoveNumFrames)
	tprev := float64(g.pinkCharacter.specialMoveCurrentFrame-1) / float64(g.pinkCharacter.specialMoveNumFrames)
	frameHorizontalMove := util.SmoothStop6(t) - util.SmoothStop6(tprev)
	if g.pinkCharacter.facing == right {
		g.pinkCharacter.x += 0.75 * frameHorizontalMove
	} else {
		g.pinkCharacter.x -= 0.75 * frameHorizontalMove
	}
	// vertical move
	firstSplit := g.pinkCharacter.specialMoveNumFrames / 3
	secondSplit := g.pinkCharacter.specialMoveNumFrames - firstSplit
	if g.pinkCharacter.specialMoveCurrentFrame <= firstSplit {
		t := float64(g.pinkCharacter.specialMoveCurrentFrame) / float64(firstSplit)
		tprev := float64(g.pinkCharacter.specialMoveCurrentFrame-1) / float64(firstSplit)
		frameVerticalMove := util.SmoothStop6(t) - util.SmoothStop6(tprev)
		g.pinkCharacter.y -= 0.25 * frameVerticalMove
	} else {
		t := float64(g.pinkCharacter.specialMoveCurrentFrame-firstSplit) / float64(secondSplit)
		tprev := float64(g.pinkCharacter.specialMoveCurrentFrame-1-firstSplit) / float64(secondSplit)
		frameVerticalMove := util.SmoothStart3(t) - util.SmoothStart3(tprev)
		g.pinkCharacter.y += 1.25 * frameVerticalMove
	}
	// end of move
	if g.pinkCharacter.specialMoveCurrentFrame >= g.pinkCharacter.specialMoveNumFrames {
		if g.state != tuto1 {
			g.state = playingPink
		} else {
			g.tutoStep = 0
		}
		g.pinkCharacter.specialMoveCurrentFrame = 0
		g.pinkCharacter.state = idle
		g.pinkCharacter.y = math.Round(g.pinkCharacter.y)
	}
}

func (g *Game) performPinkSpecialMoveUp() {
	g.pinkCharacter.specialMoveCurrentFrame++
	if g.pinkCharacter.facing == right {
		g.pinkCharacter.x += 0.75 / float64(g.pinkCharacter.specialMoveNumFrames)
	} else {
		g.pinkCharacter.x -= 0.75 / float64(g.pinkCharacter.specialMoveNumFrames)
	}
	// vertical move
	firstSplit := g.pinkCharacter.specialMoveNumFrames / 3
	secondSplit := g.pinkCharacter.specialMoveNumFrames - firstSplit
	if g.pinkCharacter.specialMoveCurrentFrame <= firstSplit {
		t := float64(g.pinkCharacter.specialMoveCurrentFrame) / float64(firstSplit)
		tprev := float64(g.pinkCharacter.specialMoveCurrentFrame-1) / float64(firstSplit)
		frameVerticalMove := util.SmoothStop3(t) - util.SmoothStop3(tprev)
		g.pinkCharacter.y -= 1.25 * frameVerticalMove
	} else {
		t := float64(g.pinkCharacter.specialMoveCurrentFrame-firstSplit) / float64(secondSplit)
		tprev := float64(g.pinkCharacter.specialMoveCurrentFrame-1-firstSplit) / float64(secondSplit)
		frameVerticalMove := util.SmoothStart3(t) - util.SmoothStart3(tprev)
		g.pinkCharacter.y += 0.25 * frameVerticalMove
	}
	if g.pinkCharacter.specialMoveCurrentFrame >= g.pinkCharacter.specialMoveNumFrames {
		if g.state != tuto1 {
			g.state = playingPink
		} else {
			g.tutoStep = 1
		}
		g.pinkCharacter.specialMoveCurrentFrame = 0
		g.pinkCharacter.state = idle
		g.pinkCharacter.y = math.Round(g.pinkCharacter.y)
	}
}
