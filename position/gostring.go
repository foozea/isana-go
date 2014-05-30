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

	"fmt"
)

type GoString struct {
	ID    int
	Stone Stone
	Value Bitboard
}

type GoStringMap [361]int
type GoStringIdentifier [361](*GoString)

func (pos *Position) GetString(vx Vertex) (id int, str *GoString) {
	if !vx.IsValid() {
		return -1, nil
	}
	id = pos.GoStringMap[vx.Index]
	if id <= 0 {
		return 0, nil
	}
	str = pos.GoStrings[id]
	return
}

func (pos *Position) CreateString(s Stone, vx Vertex) int {
	initialize := func(vertex Vertex) {
		id, g := pos.GetString(vertex)
		if g == nil || g.Stone != s {
			return
		}
		for i, v := range pos.GoStringMap {
			if v == id {
				pos.GoStringMap[i] = 0
			}
		}
	}
	for i := 1; i < len(pos.GoStrings); i++ {
		if pos.GoStrings[i] == nil {
			initialize(vx.Up())
			initialize(vx.Down())
			initialize(vx.Left())
			initialize(vx.Right())
			value := Bitboard{}
			if pos.classification(vx, s, i, &value) == 0 {
				g := GoString{i, s, value}
				pos.GoStrings[i] = &g
			}
			return i
		}
	}
	panic("overflow: number of go-strings is over the limit.")
}

func (pos *Position) UpdateStrings() {
	for i, _ := range pos.GoStringMap {
		pos.GoStringMap[i] = 0
	}
	for i, _ := range pos.GoStrings {
		pos.GoStrings[i] = nil
	}
	stringId := 1
	for i := 0; i < pos.Size.Capacity(); i++ {
		v := Vertex{i, pos.Size}
		s := pos.GetStone(v)
		if s != Empty {
			value := Bitboard{}
			if pos.classification(v, s, stringId, &value) == 0 {
				g := GoString{stringId, s, value}
				pos.GoStrings[stringId] = &g
				stringId++
			}
		}
	}
}

func (pos *Position) classification(
	v Vertex,
	s Stone,
	id int,
	value *Bitboard) int {
	///
	if v.IsValid() && pos.GoStringMap[v.Index] > 0 {
		return pos.GoStringMap[v.Index]
	}
	if pos.GetStone(v) != s {
		return -1
	}
	value.SetBit(v.Index)
	pos.GoStringMap[v.Index] = id
	pos.classification(v.Up(), s, id, value)
	pos.classification(v.Down(), s, id, value)
	pos.classification(v.Left(), s, id, value)
	pos.classification(v.Right(), s, id, value)
	return 0
}

func (pos *Position) GoStringDump() {
	ls := int(pos.Size)
	files := "ABCDEFGHJKLMNOPQRSTUVWXYZ"
	// Header
	fmt.Printf("\n")
	for i := 0; i < ls; i++ {
		fmt.Printf("%2v ", string(files[i]))
	}
	fmt.Printf("\n")
	// Body
	for i := 0; i < ls; i++ {
		for j := 0; j < ls; j++ {
			v := pos.GoStringMap[ls*(ls-1)-(ls*i)+j]
			if v == 0 {
				fmt.Printf(" . ")
			} else {
				fmt.Printf("%2v ", v)
			}
			if (j+1)%ls == 0 {
				fmt.Printf(" %v\n", ls-i)
			}
		}
	}
	fmt.Printf("\n")
}
