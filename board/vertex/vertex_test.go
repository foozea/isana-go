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

package vertex

import (
	"testing"

	. "github.com/foozea/isana/board/size"
)

func TestIsValid(t *testing.T) {
	var actual, expected bool
	const msg string = "IsValid / Miss detected expected : %v, but %v"
	actual = Vertex{-1, B9x9}.IsValid()
	expected = false
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = Vertex{5, B9x9}.IsValid()
	expected = true
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = Vertex{83, B9x9}.IsValid()
	expected = false
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}

func TestStringToVertex(t *testing.T) {
	var actual, expected Vertex
	const msg string = "StringToVertex / couldn't convert from string. expected : %v, but %v"
	actual = StringToVertex("A1", B9x9)
	expected = Vertex{0, B9x9}
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = StringToVertex("E3", B9x9)
	expected = Vertex{22, B9x9}
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = StringToVertex("J5", B9x9)
	expected = Vertex{44, B9x9}
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = StringToVertex("J19", B9x9)
	expected = Outbound
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}

func TestString(t *testing.T) {
	var actual, expected string
	const msg string = "String /could't convert to string. expected : %v, but %v"
	actual = StringToVertex("E3", B9x9).String()
	expected = "E3"
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = StringToVertex("J7", B9x9).String()
	expected = "J7"
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = Outbound.String()
	expected = "PASS"
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}

func TestUp(t *testing.T) {
	var actual, expected Vertex
	const msg string = "Up / the vertex didn't move to up expected : %v, but %v"
	actual = StringToVertex("E2", B9x9).Up()
	expected = StringToVertex("E3", B9x9)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}

func TestDown(t *testing.T) {
	var actual, expected Vertex
	const msg string = "Down / the vertex didn't move to down. expected : %v, but %v"
	actual = StringToVertex("E2", B9x9).Down()
	expected = StringToVertex("E1", B9x9)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}

func TestLeft(t *testing.T) {
	var actual, expected Vertex
	const msg string = "Left /  the vertex didn't move to left. expected : %v, but %v"
	actual = StringToVertex("E2", B9x9).Left()
	expected = StringToVertex("D2", B9x9)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}

func TestRight(t *testing.T) {
	var actual, expected Vertex
	const msg string = "Right / the vertex didn't move to right. expected : %v, but %v"
	actual = StringToVertex("F2", B9x9).Right()
	expected = StringToVertex("G2", B9x9)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}
