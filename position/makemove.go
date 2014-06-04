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
	. "github.com/foozea/isana/board/stone"
	. "github.com/foozea/isana/board/vertex"
	. "github.com/foozea/isana/position/move"
)

// Take stones these are `DEAD`.
func (pos *Position) TakeStone(stone Stone, vx Vertex) int {
	id, g := pos.GetString(vx)
	if g == nil || g.Stone != stone || pos.CountLiberty(id, g) != 0 {
		return 0
	}
	prisoners := g.Value.CountBit()
	if stone == Black {
		pos.blacks = Xor(pos.blacks, g.Value)
	} else if stone == White {
		pos.whites = Xor(pos.whites, g.Value)
	}
	pos.GoStrings[id] = nil
	for i, v := range pos.GoStringMap {
		if v == g.ID {
			pos.GoStringMap[i] = 0
		}
	}
	return prisoners
}

// Make pseudo move and validate it.
// if invalid move (it is illegal, or suicide), returns false as ok-code.
func (pos *Position) PseudoMove(mv *Move) (next *Position, ok bool) {
	if !pos.isLegalMove(mv) {
		return nil, false
	}
	test := CopyPosition(pos)
	test.SetStone(mv.Stone, mv.Vertex)
	test.CreateString(mv.Stone, mv.Vertex)
	if test.isSuicideMove(mv, 0) {
		return nil, false
	}
	return &test, true
}

// Make pseudo move and validate it.
// if invalid move (it is illegal, suicide or filling own eye),
// returns false as ok-code.
func (pos *Position) PseudoMoveStrict(mv *Move) (next *Position, ok bool) {
	if !pos.isLegalMove(mv) || pos.isFillEyeMove(mv) {
		return nil, false
	}
	test := CopyPosition(pos)
	test.SetStone(mv.Stone, mv.Vertex)
	test.CreateString(mv.Stone, mv.Vertex)
	// strict mode
	if test.isSuicideMove(mv, 1) {
		return nil, false
	}
	return &test, true
}

// Make a move. If the neighbours are DEAD, take it and add them to prison.
// if it is Ko, memory it.
func (pos *Position) FixMove(move *Move) {
	pos.SetStone(move.Stone, move.Vertex)
	pos.CreateString(move.Stone, move.Vertex)
	opp := move.Stone.Opposite()
	v := move.Vertex

	dir := []Vertex{v.Up(), v.Down(), v.Left(), v.Right()}
	ss := []Stone{pos.GetStone(dir[0]), pos.GetStone(dir[1]), pos.GetStone(dir[2]), pos.GetStone(dir[3])}

	prisonersCount := 0
	for _, v := range dir {
		prisonersCount += pos.TakeStone(opp, v)
	}
	if move.Stone == Black {
		pos.BlackPrison += prisonersCount
		pos.WhitePrison += pos.TakeStone(move.Stone, v)
	} else {
		pos.WhitePrison += prisonersCount
		pos.BlackPrison += pos.TakeStone(move.Stone, v)
	}

	if (ss[0] == opp || ss[0] == Wall) &&
		(ss[1] == opp || ss[1] == Wall) &&
		(ss[2] == opp || ss[2] == Wall) &&
		(ss[3] == opp || ss[3] == Wall) {
		if prisonersCount == 1 {
			for i, v := range ss {
				if v == Empty {
					pos.Ko = dir[i]
					pos.KoStone = move.Stone
					break
				}
			}
			return
		}
	}
	pos.Ko, pos.KoStone = Outbound, Empty
}
