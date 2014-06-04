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

type HashboardType map[int]uint64

var Hashboard HashboardType

func init() {
	Seed(Now().UTC().UnixNano())
	Hashboard = HashboardType{
		6: 3201594604936722833, 13: 2937994859130440336, 9: 6373431332469792527,
		17: 14473860272246989144, 25: 18281815577875503594, 30: 10061258553119313312,
		33: 3286497388246993722, 1: 3985502496055295820, 2: 7006577664179485562,
		5: 10892363896395695575, 10: 12801914950157331807, 18: 16392973390174980921,
		22: 15624686782962010408, 14: 16923473909756009598, 21: 2429104740398281672,
		26: 13756929997955362323, 29: 14123451814094485062, 34: 6639239958766330982}
}

func (mb Miniboard) Hashcode() uint64 {
	code := uint64(0x0)
	for i, v := range mb {
		code ^= Hashboard[i<<2|int(v)]
	}
	return code
}

func (mb Miniboard) Rotate90() Miniboard {
	return Miniboard{mb[6], mb[3], mb[0], mb[7], mb[4], mb[1], mb[8], mb[5], mb[2]}
}

func (mb Miniboard) Rotate180() Miniboard {
	return mb.Rotate90().Rotate90()
}

func (mb Miniboard) Rotate270() Miniboard {
	return mb.Rotate180().Rotate90()
}

func (mb Miniboard) Mirror() Miniboard {
	return Miniboard{mb[2], mb[1], mb[0], mb[5], mb[4], mb[3], mb[8], mb[7], mb[6]}
}

func (mb Miniboard) Invert() Miniboard {
	ret := Miniboard{}
	for i, v := range mb {
		if v == Wall {
			ret[i] = Wall
		} else {
			ret[i] = v.Opposite()
		}
	}
	return ret
}

func (mb Miniboard) AllAddAs(cat Category, patterns map[uint64]Category) {
	patterns[mb.Hashcode()] = cat
	patterns[mb.Rotate90().Hashcode()] = cat
	patterns[mb.Rotate180().Hashcode()] = cat
	patterns[mb.Rotate270().Hashcode()] = cat
	patterns[mb.Mirror().Hashcode()] = cat
	patterns[mb.Mirror().Rotate90().Hashcode()] = cat
	patterns[mb.Mirror().Rotate180().Hashcode()] = cat
	patterns[mb.Mirror().Rotate270().Hashcode()] = cat
}

// hash code for 3x3 patterns
func createHash() HashboardType {
	hash := make(HashboardType, 0)
	for i := 0; i < 9; i++ {
		hash[i<<2|int(Black)] = uint64(Uint32())<<32 | uint64(Uint32())
		hash[i<<2|int(White)] = uint64(Uint32())<<32 | uint64(Uint32())
	}
	return hash
}
