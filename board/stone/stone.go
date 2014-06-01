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
	"strings"
)

type Stone uint

const (
	Empty Stone = 0
	Black Stone = 1
	White Stone = 2
	Wall  Stone = 3
)

func (s Stone) Opposite() Stone {
	switch s {
	case Black:
		return White
	case White:
		return Black
	}
	return Empty
}

func StringToStone(str string) Stone {
	str = strings.ToUpper(str)
	if str == "BLACK" || str == "B" {
		return Black
	}
	if str == "WHITE" || str == "W" {
		return White
	}
	if str == "RESIGN" {
		return Empty
	}
	return Wall
}

func (s Stone) String() string {
	switch s {
	case Empty:
		return "."
	case Black:
		return "X"
	case White:
		return "O"
	}
	return " "
}
