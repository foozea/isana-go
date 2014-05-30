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

package position

import (
	. "code.isana.io/isana/board/bitboard"
	. "code.isana.io/isana/board/stone"
	. "code.isana.io/isana/board/vertex"
)

func (pos *Position) GetStone(vx Vertex) Stone {
	if !vx.IsValid() {
		return Wall
	}
	if pos.blacks.GetBit(vx.Index) == 1 {
		return Black
	} else if pos.whites.GetBit(vx.Index) == 1 {
		return White
	} else {
		return Empty
	}
}

func (pos *Position) SetStone(stone Stone, vx Vertex) bool {
	if !vx.IsValid() || stone == Wall {
		return false
	}
	if stone == Black {
		pos.blacks.SetBit(vx.Index)
	} else if stone == White {
		pos.whites.SetBit(vx.Index)
	} else {
		pos.blacks.ClearBit(vx.Index)
		pos.whites.ClearBit(vx.Index)
	}
	return true
}

func (pos *Position) Empties() [](*Vertex) {
	vs := make([](*Vertex), 0)
	bits := Or(pos.blacks, pos.whites)
	for i := 0; i < pos.Size.Capacity(); i++ {
		if bits.GetBit(i) == 0 {
			v := Vertex{i, pos.Size}
			vs = append(vs, &v)
		}
	}
	return vs
}