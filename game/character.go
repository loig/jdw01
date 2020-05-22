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

type character struct {
	x                       float64
	y                       float64
	speed                   float64
	state                   characterState
	previousState           characterState
	facing                  side
	animationFrame          int
	animationStep           int
	idleFrames              []int
	moveFrames              []int
	specialMoveFrames       []int
	specialMoveNumFrames    int
	specialMoveCurrentFrame int
	strikeFrames            []int
	strikeNumFrames         int
	strikeCurrentFrame      int
}

type characterState int

const (
	idle characterState = iota
	move
	specialMove
	strike
)

type side int

const (
	right side = iota
	left
)
