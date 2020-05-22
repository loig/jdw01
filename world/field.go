// Package world implements world representation and generation
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
package world

// FieldTile defines one tile of the field
type FieldTile struct {
	Tile       FieldType
	Decoration FieldType
	Destructed FieldType
}

// FieldType defines the type of a (part of a) tile
type FieldType struct {
	kind  fieldKind
	LookX int
	LookY int
}

// FieldKind are the different kind of possible field types
type fieldKind int

// All the possible field kinds
const (
	floor fieldKind = iota
	traversableWall
	destroyableWall
	backgroundWall
	wall
	lader
	nothing
)

// IsBackgroundField states which tiles can be behind characters
func IsBackgroundField(field FieldTile) bool {
	switch field.Tile.kind {
	case nothing, backgroundWall, lader:
		return true
	}
	return false
}

// IsLaderField states which tiles are compatible with the
// white character special move
func IsLaderField(field FieldTile) bool {
	return field.Tile.kind == lader || field.Decoration.kind == lader
}

// IsFloorField states on which tiles the characters can walk
func IsFloorField(field FieldTile) bool {
	switch field.Tile.kind {
	case floor, destroyableWall, traversableWall, wall:
		return true
	}
	return false
}

// IsTraversableField states which tiles are compatible with the
// blue character special move
func IsTraversableField(field FieldTile) bool {
	switch field.Tile.kind {
	case traversableWall:
		return true
	}
	return false
}

// IsBreakableField states which tiles can be broken when striked
// by the characters
func IsBreakableField(field FieldTile) bool {
	switch field.Tile.kind {
	case destroyableWall:
		return true
	}
	return false
}
