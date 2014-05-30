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

package move

import (
	"testing"

	. "code.isana.io/isana/board/size"
	. "code.isana.io/isana/board/stone"
	. "code.isana.io/isana/board/vertex"
)

func TestCreateMove(t *testing.T) {
	var actual, expected Move
	const msg string = "CreateMove / couldn't create valid move. expected : %v, but %v"
	actual = *CreateMove(Black, Vertex{5, B9x9})
	expected = Move{Black, Vertex{5, B9x9}, 0.0, 0.0}
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}

func TestString(t *testing.T) {
	var actual, expected string
	const msg string = "String / convert faild. expected : %v, but %v"
	actual = CreateMove(Black, StringToVertex("E4", B9x9)).String()
	expected = "E4"
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}
