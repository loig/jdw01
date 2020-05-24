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
	_ "image/png" //for ebitenutil.NewImageFromFile

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var blueCharacterImage *ebiten.Image
var whiteCharacterImage *ebiten.Image
var pinkCharacterImage *ebiten.Image
var tilesImage *ebiten.Image
var dayBackgroundImage *ebiten.Image
var buttonRightImage *ebiten.Image
var buttonLeftImage *ebiten.Image
var buttonXImage *ebiten.Image
var buttonUpImage *ebiten.Image
var buttonDownImage *ebiten.Image
var buttonLImage *ebiten.Image
var buttonRImage *ebiten.Image

func loadAssets() (err error) {

	blueCharacterImage, _, err = ebitenutil.NewImageFromFile("assets/dude.png", ebiten.FilterDefault)
	if err != nil {
		return err
	}

	whiteCharacterImage, _, err = ebitenutil.NewImageFromFile("assets/owlet.png", ebiten.FilterDefault)
	if err != nil {
		return err
	}

	pinkCharacterImage, _, err = ebitenutil.NewImageFromFile("assets/pink.png", ebiten.FilterDefault)
	if err != nil {
		return err
	}

	tilesImage, _, err = ebitenutil.NewImageFromFile("assets/tiles.png", ebiten.FilterDefault)
	if err != nil {
		return err
	}

	dayBackgroundImage, _, err = ebitenutil.NewImageFromFile("assets/daybackground.png", ebiten.FilterDefault)
	if err != nil {
		return err
	}

	buttonRightImage, _, err = ebitenutil.NewImageFromFile("assets/rightbutton.png", ebiten.FilterDefault)
	if err != nil {
		return err
	}

	buttonLeftImage, _, err = ebitenutil.NewImageFromFile("assets/leftbutton.png", ebiten.FilterDefault)
	if err != nil {
		return err
	}

	buttonXImage, _, err = ebitenutil.NewImageFromFile("assets/xbutton.png", ebiten.FilterDefault)
	if err != nil {
		return err
	}

	buttonUpImage, _, err = ebitenutil.NewImageFromFile("assets/upbutton.png", ebiten.FilterDefault)
	if err != nil {
		return err
	}

	buttonDownImage, _, err = ebitenutil.NewImageFromFile("assets/downbutton.png", ebiten.FilterDefault)
	if err != nil {
		return err
	}

	buttonLImage, _, err = ebitenutil.NewImageFromFile("assets/lbutton.png", ebiten.FilterDefault)
	if err != nil {
		return err
	}

	buttonRImage, _, err = ebitenutil.NewImageFromFile("assets/rbutton.png", ebiten.FilterDefault)
	if err != nil {
		return err
	}

	return nil
}
