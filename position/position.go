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
	. "github.com/foozea/isana/board/bitboard"
	. "github.com/foozea/isana/board/size"
	. "github.com/foozea/isana/board/stone"
	. "github.com/foozea/isana/board/vertex"
	. "github.com/foozea/isana/position/move"

	"fmt"
)

type Position struct {
	blacks Bitboard
	whites Bitboard
	Size   BoardSize
	///
	Ko      Vertex
	KoStone Stone
	///
	BlackPrison int
	WhitePrison int
	///
	GoStringMap GoStringMap
	GoStrings   GoStringIdentifier
	///
	Games         float64
	Moves         []Move
	ProbDencities [361]int
	TotalProbs    int
	SubTotalProbs [19]int
}

func CreatePosition(size BoardSize) Position {
	return Position{
		Bitboard{},
		Bitboard{},
		size,
		Outbound,
		Empty,
		0, 0,
		GoStringMap{},
		GoStringIdentifier{},
		0.0,
		make([]Move, 0),
		[361]int{},
		0, [19]int{}}
}

func CopyPosition(pos *Position) Position {
	copied := Position{
		pos.blacks,
		pos.whites,
		pos.Size,
		pos.Ko,
		pos.KoStone,
		pos.BlackPrison,
		pos.WhitePrison,
		pos.GoStringMap,
		pos.GoStrings,
		0.0,
		make([]Move, 0),
		[361]int{},
		0, [19]int{}}

	return copied
}

func (pos *Position) CountStringLiberty(id int, g *GoString) int {
	// outbound of the board
	if id < 0 {
		return 0
	}
	// empty
	if g == nil {
		return 1
	}
	n := Xor(Or(
		LeftShift(g.Value, pos.Size),
		RightShift(g.Value, pos.Size),
		UpShift(g.Value, pos.Size),
		DownShift(g.Value, pos.Size),
		g.Value), g.Value)

	var bits Bitboard
	if g.Stone == Black {
		bits = pos.whites
	} else if g.Stone == White {
		bits = pos.blacks
	}

	m := And(n, bits)
	return n.CountBit() - m.CountBit()
}

func (pos *Position) Score(stone Stone, komi float64) float64 {
	score := 0.0
	for id, v := range pos.GoStrings {
		if v != nil {
			delta := 0.0
			if pos.CountStringLiberty(id, v) > 1 {
				delta = float64(v.Value.CountBit())
			}
			if v.Stone == Black {
				score += delta
			} else if v.Stone == White {
				score -= delta
			}
		}
	}
	for _, v := range pos.Empties() {
		up, down, left, right := v.Up(), v.Down(), v.Left(), v.Right()
		bn := pos.blacks.GetBit(up.Index) +
			pos.blacks.GetBit(down.Index) +
			pos.blacks.GetBit(left.Index) +
			pos.blacks.GetBit(right.Index)
		wn := pos.whites.GetBit(up.Index) +
			pos.whites.GetBit(down.Index) +
			pos.whites.GetBit(left.Index) +
			pos.whites.GetBit(right.Index)
		if bn != 0 && wn == 0 {
			score++
		} else if bn == 0 && wn != 0 {
			score--
		}
	}
	score -= komi
	score += float64(pos.BlackPrison - pos.WhitePrison)
	win := 0.0
	if score > 0.0 {
		win = 1.0
	}
	if stone == White {
		win = -win
	}
	return win
}

func (pos *Position) Dump() {
	ls := int(pos.Size)
	files := "ABCDEFGHJKLMNOPQRSTUVWXYZ"
	stones := make([]Stone, pos.Size.Capacity())
	for i := 0; i < pos.Size.Capacity(); i++ {
		stones[i] = pos.GetStone(Vertex{i, pos.Size})
	}
	// Header
	fmt.Printf("\n")
	for i := 0; i < ls; i++ {
		fmt.Printf("%2v ", string(files[i]))
	}
	fmt.Printf("\n")
	// Body
	for i := 0; i < ls; i++ {
		n := len(stones)
		sts := stones[n-ls*(i+1) : n-ls*i]
		for j, v := range sts {
			v.Dump()
			fmt.Printf(" ")
			if (j+1)%ls == 0 {
				fmt.Printf(" %v\n", ls-i)
			}
		}
	}
	fmt.Printf("\n")
}
