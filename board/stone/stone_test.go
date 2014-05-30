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

package stone

import (
	"fmt"
	"testing"
)

func init() {
	fmt.Println("Dump...")
	Black.Dump()
	White.Dump()
	Empty.Dump()
	Wall.Dump()
	fmt.Println("")
}

func TestOpposite(t *testing.T) {
	var expected, actual Stone
	const msg string = "Opposite / coudl'nt get opposite stone. expected : %v, but %v"
	expected = White
	actual = Black.Opposite()
	if expected != actual {
		t.Errorf(msg, expected, actual)
	}
	expected = Black
	actual = White.Opposite()
	if expected != actual {
		t.Errorf(msg, expected, actual)
	}
	expected = Empty
	actual = Empty.Opposite()
	if expected != actual {
		t.Errorf(msg, expected, actual)
	}
	expected = Empty
	actual = Wall.Opposite()
	if expected != actual {
		t.Errorf(msg, expected, actual)
	}
}

func TestStringToStone(t *testing.T) {
	var expected, actual Stone
	const msg string = "StringToStone / couldn't convert from string. expected : %v, but %v"
	actual = StringToStone("Black")
	expected = Black
	if expected != actual {
		t.Errorf(msg, expected, actual)
	}
	actual = StringToStone("b")
	expected = Black
	if expected != actual {
		t.Errorf(msg, expected, actual)
	}
	actual = StringToStone("WHITE")
	expected = White
	if expected != actual {
		t.Errorf(msg, expected, actual)
	}
	actual = StringToStone("W")
	expected = White
	if expected != actual {
		t.Errorf(msg, expected, actual)
	}
	actual = StringToStone("Resign")
	expected = Empty
	if expected != actual {
		t.Errorf(msg, expected, actual)
	}
	actual = StringToStone("AAA")
	expected = Wall
	if expected != actual {
		t.Errorf(msg, expected, actual)
	}
}
