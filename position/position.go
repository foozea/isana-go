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
	"fmt"

	. "github.com/foozea/isana/board/bitboard"
	. "github.com/foozea/isana/board/size"
	. "github.com/foozea/isana/board/stone"
	. "github.com/foozea/isana/board/vertex"
	. "github.com/foozea/isana/hashing"
	. "github.com/foozea/isana/position/move"
)

// Go-Board struct.
// This is also keep various information for tree serch.
type Position struct {
	hash uint64
	///
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
	Games         int32
	RaveGames     int32
	Moves         []Move
	ProbDencities [361]int
	TotalProbs    int
	SubTotalProbs [19]int
}

// Returns empty Go-board.
func CreatePosition(size BoardSize) Position {
	return Position{
		0x0,
		Bitboard{},
		Bitboard{},
		size,
		Outbound,
		Empty,
		0, 0,
		GoStringMap{},
		GoStringIdentifier{},
		0, 0,
		make([]Move, 0),
		[361]int{},
		0, [19]int{}}
}

// Copies the Go-board.
func CopyPosition(pos *Position) Position {
	copied := Position{
		pos.hash,
		pos.blacks,
		pos.whites,
		pos.Size,
		pos.Ko,
		pos.KoStone,
		pos.BlackPrison,
		pos.WhitePrison,
		pos.GoStringMap,
		pos.GoStrings,
		0, 0,
		make([]Move, 0),
		pos.ProbDencities,
		pos.TotalProbs,
		pos.SubTotalProbs}

	return copied
}

// Gets the stone of the vertex.
// It returns Wall for invalid vertex, and return s empty
// if no stone is set at the vertex.
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

// Set the stone to the vertex.
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

// Returns a slice of empty vertex
func (pos *Position) Empties() []Vertex {
	vs := make([]Vertex, 0)
	bits := Or(pos.blacks, pos.whites)
	for i := 0; i < pos.Size.Capacity(); i++ {
		if bits.GetBit(i) == 0 {
			vs = append(vs, Vertex{i, pos.Size})
		}
	}
	return vs
}

// Counts liberty number of the GoString.
func (pos *Position) CountLiberty(id int, g *GoString) int {
	// outbound of the board
	if id < 0 {
		return 0
	}
	// empty
	if g == nil {
		return 1
	}
	n := Xor(Or(
		Left(g.Value, pos.Size), Right(g.Value, pos.Size),
		Up(g.Value, pos.Size), Down(g.Value, pos.Size),
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

// Counts the score of the position.
// - black wins => 1
//   black loses => 0
//   white wins => 0
//   white loses => -1
func (pos *Position) Score(stone Stone, komi float64) float64 {
	score := 0
	for id, v := range pos.GoStrings {
		if v != nil {
			delta := 0
			if pos.CountLiberty(id, v) > 1 {
				delta = v.Value.CountBit()
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
		///
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
	score += pos.BlackPrison - pos.WhitePrison
	if float64(score)-komi > 0.0 {
		if stone == White {
			return -1.0
		}
		return 1.0
	}
	return 0
}

// trim 3x3 square and returns the hash code.
func (pos *Position) SquaredHash3x3(v Vertex) uint64 {
	if !v.IsValid() {
		return 0
	}
	delta := func(arg Vertex, i int) uint64 {
		s := pos.GetStone(arg)
		return Hashboard[i<<2|int(s)]
	}
	// treats starting point as index '5'.
	return delta(v.Down().Left(), 1) ^ // 1
		delta(v.Down(), 2) ^ // 2
		delta(v.Down().Right(), 3) ^ // 3
		delta(v.Left(), 4) ^ // 4
		delta(v, 5) ^ // 5
		delta(v.Right(), 6) ^ // 6
		delta(v.Up().Left(), 7) ^ // 7
		delta(v.Up(), 8) ^ // 8
		delta(v.Up().Right(), 9) // 9
}

// trim 3x3 square and returns the hash code.
func (pos *Position) SquaredHash3x2(v Vertex) (_1 uint64, _2 uint64) {
	if !v.IsValid() {
		return 0, 0
	}
	delta := func(arg Vertex, i int) uint64 {
		s := pos.GetStone(arg)
		return Hashboard[i<<2|int(s)]
	}
	_1 = delta(v.Down().Left(), 1) ^ // 1
		delta(v.Down(), 2) ^ // 2
		delta(v.Down().Right(), 3) ^ // 3
		delta(v.Left(), 4) ^ // 4
		delta(v, 5) ^ // 5
		delta(v.Right(), 6) // 6
	_2 = delta(v.Left(), 4) ^ // 4
		delta(v, 5) ^ // 5
		delta(v.Right(), 6) ^ // 6
		delta(v.Up().Left(), 7) ^ // 7
		delta(v.Up(), 8) ^ // 8
		delta(v.Up().Right(), 9) // 9
	return _1, _2
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
			fmt.Printf(" %v ", v.String())
			if (j+1)%ls == 0 {
				fmt.Printf(" %v\n", ls-i)
			}
		}
	}
	fmt.Printf("\n")
}
