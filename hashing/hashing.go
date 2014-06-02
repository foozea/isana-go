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

package hashing

import (
	. "math/rand"
	. "time"

	. "github.com/foozea/isana/board/stone"
)

type DeltaHashType map[int]uint64

var DeltaHash DeltaHashType

func init() {
	Seed(Now().UTC().UnixNano())
	DeltaHash = createHash()
}

// hash code for 3x3 patterns
func createHash() DeltaHashType {
	deltaHash := make(DeltaHashType, 0)
	for i := 0; i < 9; i++ {
		deltaHash[i<<2|int(Black)] = uint64(Uint32())<<32 | uint64(Uint32())
		deltaHash[i<<2|int(White)] = uint64(Uint32())<<32 | uint64(Uint32())
	}
	return deltaHash
}
