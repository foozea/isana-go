/*
  Isana, a software for the game of Go
  Copyright (C) 2014 Tetsuo FUJII

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package control

import (
	"testing"

	. "github.com/foozea/isana/board/size"
	. "github.com/foozea/isana/board/stone"
)

var (
	state GameState
)

func init() {
	state = CreateDefaultGameState()
}

func TestCreate(t *testing.T) {
	var actual, expected GameState
	const msg string = "CreateDefaultGameState / failed to create default state. expected : %v, but %v"
	actual = state
	expected = GameState{B9x9, nil, Empty, 0.0, TimeSettings{60, 600, 25}}
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}

func TestCurrentStoneIs(t *testing.T) {
	var actual, expected Stone
	const msg string = "CurrentStoneIs / failed to get current stone. expected : %v, but %v"
	actual = state.CurrentStoneIs()
	expected = Empty
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	state.Turn = Black
	actual = state.CurrentStoneIs()
	expected = Black
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}
