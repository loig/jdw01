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
	"io/ioutil"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type audioManagerInfo struct {
	audioContext *audio.Context
	soundPlayer  *audio.Player
	currentSound soundType
}

type soundType int

const (
	moveSound soundType = iota
	whitemoveSound
	pinkSound
	blueSound
	whiteSound
	noSound
	strikeSound
	pinkstrikeSound
	whitestrikeSound
)

var (
	moveSoundBytes        []byte
	pinkSoundBytes        []byte
	blueSoundBytes        []byte
	strikeSoundBytes      []byte
	pinkstrikeSoundBytes  []byte
	whitestrikeSoundBytes []byte
	whiteSoundBytes       []byte
	whitemoveSoundBytes   []byte
)

func (g *Game) initSound() (err error) {
	g.audioManager.audioContext, err = audio.NewContext(44100)

	soundFile, err := ebitenutil.OpenFile("assets/move.mp3")
	if err != nil {
		return err
	}
	sound, err := mp3.Decode(g.audioManager.audioContext, soundFile)
	if err != nil {
		return err
	}
	moveSoundBytes, err = ioutil.ReadAll(sound)
	if err != nil {
		return err
	}

	soundFile, err = ebitenutil.OpenFile("assets/pinkSpecial.mp3")
	if err != nil {
		return err
	}
	sound, err = mp3.Decode(g.audioManager.audioContext, soundFile)
	if err != nil {
		return err
	}
	pinkSoundBytes, err = ioutil.ReadAll(sound)
	if err != nil {
		return err
	}

	soundFile, err = ebitenutil.OpenFile("assets/blueSpecial.mp3")
	if err != nil {
		return err
	}
	sound, err = mp3.Decode(g.audioManager.audioContext, soundFile)
	if err != nil {
		return err
	}
	blueSoundBytes, err = ioutil.ReadAll(sound)
	if err != nil {
		return err
	}

	soundFile, err = ebitenutil.OpenFile("assets/strike.mp3")
	if err != nil {
		return err
	}
	sound, err = mp3.Decode(g.audioManager.audioContext, soundFile)
	if err != nil {
		return err
	}
	strikeSoundBytes, err = ioutil.ReadAll(sound)
	if err != nil {
		return err
	}

	soundFile, err = ebitenutil.OpenFile("assets/pinkstrike.mp3")
	if err != nil {
		return err
	}
	sound, err = mp3.Decode(g.audioManager.audioContext, soundFile)
	if err != nil {
		return err
	}
	pinkstrikeSoundBytes, err = ioutil.ReadAll(sound)
	if err != nil {
		return err
	}

	soundFile, err = ebitenutil.OpenFile("assets/whitestrike.mp3")
	if err != nil {
		return err
	}
	sound, err = mp3.Decode(g.audioManager.audioContext, soundFile)
	if err != nil {
		return err
	}
	whitestrikeSoundBytes, err = ioutil.ReadAll(sound)
	if err != nil {
		return err
	}

	soundFile, err = ebitenutil.OpenFile("assets/lader.mp3")
	if err != nil {
		return err
	}
	sound, err = mp3.Decode(g.audioManager.audioContext, soundFile)
	if err != nil {
		return err
	}
	whiteSoundBytes, err = ioutil.ReadAll(sound)
	if err != nil {
		return err
	}

	soundFile, err = ebitenutil.OpenFile("assets/whitemove.mp3")
	if err != nil {
		return err
	}
	sound, err = mp3.Decode(g.audioManager.audioContext, soundFile)
	if err != nil {
		return err
	}
	whitemoveSoundBytes, err = ioutil.ReadAll(sound)
	if err != nil {
		return err
	}

	return err
}

func (g *Game) updateSound() {

	soundToPlay := noSound
	if g.blueCharacter.state == move ||
		g.pinkCharacter.state == move {
		soundToPlay = moveSound
	} else if g.pinkCharacter.state == specialMove {
		soundToPlay = pinkSound
	} else if g.blueCharacter.state == specialMove {
		soundToPlay = blueSound
	} else if g.blueCharacter.state == strike {
		soundToPlay = strikeSound
	} else if g.whiteCharacter.state == strike {
		soundToPlay = whitestrikeSound
	} else if g.pinkCharacter.state == strike {
		soundToPlay = pinkstrikeSound
	} else if g.state == whiteSpecialMove {
		soundToPlay = whiteSound
	} else if g.whiteCharacter.state == move {
		soundToPlay = whitemoveSound
	}

	if soundToPlay != g.audioManager.currentSound {
		switch soundToPlay {
		case noSound:
			g.audioManager.soundPlayer = nil
		case moveSound:
			g.audioManager.soundPlayer, _ = audio.NewPlayerFromBytes(
				g.audioManager.audioContext, moveSoundBytes,
			)
			g.audioManager.soundPlayer.Play()
		case whitemoveSound:
			g.audioManager.soundPlayer, _ = audio.NewPlayerFromBytes(
				g.audioManager.audioContext, whitemoveSoundBytes,
			)
			g.audioManager.soundPlayer.Play()
		case whiteSound:
			g.audioManager.soundPlayer, _ = audio.NewPlayerFromBytes(
				g.audioManager.audioContext, whiteSoundBytes,
			)
			g.audioManager.soundPlayer.Play()
		case pinkSound:
			if g.audioManager.soundPlayer != nil {
				g.audioManager.soundPlayer.Pause()
			}
			g.audioManager.soundPlayer = nil
			tmpPlayer, _ := audio.NewPlayerFromBytes(
				g.audioManager.audioContext, pinkSoundBytes,
			)
			tmpPlayer.Play()
		case blueSound:
			g.audioManager.soundPlayer = nil
			tmpPlayer, _ := audio.NewPlayerFromBytes(
				g.audioManager.audioContext, blueSoundBytes,
			)
			tmpPlayer.Play()
		case strikeSound:
			g.audioManager.soundPlayer = nil
			tmpPlayer, _ := audio.NewPlayerFromBytes(
				g.audioManager.audioContext, strikeSoundBytes,
			)
			tmpPlayer.Play()
		case whitestrikeSound:
			g.audioManager.soundPlayer = nil
			tmpPlayer, _ := audio.NewPlayerFromBytes(
				g.audioManager.audioContext, whitestrikeSoundBytes,
			)
			tmpPlayer.Play()
		case pinkstrikeSound:
			g.audioManager.soundPlayer = nil
			tmpPlayer, _ := audio.NewPlayerFromBytes(
				g.audioManager.audioContext, pinkstrikeSoundBytes,
			)
			tmpPlayer.Play()
		}
		g.audioManager.currentSound = soundToPlay
		return
	}

	if g.audioManager.soundPlayer != nil && !g.audioManager.soundPlayer.IsPlaying() {
		g.audioManager.soundPlayer.Rewind()
		g.audioManager.soundPlayer.Play()
	}
}
