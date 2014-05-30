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
	. "code.isana.io/isana/board/size"
	. "code.isana.io/isana/board/stone"
	. "code.isana.io/isana/board/vertex"
	. "code.isana.io/isana/position/move"

	. "math/rand"
	. "time"
)

type MoveHashType map[Move]uint64
type ZobristHash map[uint64](*Position)

var MoveHash MoveHashType
var PositionHash ZobristHash

func init() {
	Seed(Now().UTC().UnixNano())
}

func CreateHash(size BoardSize) {
	MoveHash = make(MoveHashType, size.Capacity()*2)
	for i := 0; i < size.Capacity(); i++ {
		v := Vertex{i, size}
		black := CreateMove(Black, v)
		white := CreateMove(White, v)
		MoveHash[*black] = uint64(Uint32())<<32 | uint64(Uint32())
		MoveHash[*white] = uint64(Uint32())<<32 | uint64(Uint32())
	}
}
