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
	if !pos.isLegalMove(mv) {
		return nil, false
	}
	test := CopyPosition(pos)
	test.SetStone(mv.Stone, mv.Vertex)
	test.CreateString(mv.Stone, mv.Vertex)
	// strict mode
	if test.isSuicideMove(mv, 1) || test.isFillEyeMove(mv) {
		return nil, false
	}
	return &test, true
}

// Make a move. If the neighbours are DEAD, take it and add them to prison.
// if it is Ko, memory it.
func (pos *Position) FixMove(move *Move) {
	pos.SetStone(move.Stone, move.Vertex)
	opp, v := move.Stone.Opposite(), move.Vertex
	_ = pos.CreateString(move.Stone, move.Vertex)

	up, down, left, right :=
		v.Up(), v.Down(), v.Left(), v.Right()
	upS, downS, leftS, rightS :=
		pos.GetStone(up), pos.GetStone(down), pos.GetStone(left), pos.GetStone(right)

	prisonersCount := pos.TakeStone(opp, up)
	prisonersCount += pos.TakeStone(opp, down)
	prisonersCount += pos.TakeStone(opp, left)
	prisonersCount += pos.TakeStone(opp, right)
	suicideCount := pos.TakeStone(move.Stone, v)

	if move.Stone == Black {
		pos.BlackPrison += prisonersCount
		pos.WhitePrison += suicideCount
	} else {
		pos.WhitePrison += prisonersCount
		pos.BlackPrison += suicideCount
	}
	ko_flag := false
	if (leftS == opp || leftS == Wall) &&
		(rightS == opp || rightS == Wall) &&
		(upS == opp || upS == Wall) &&
		(downS == opp || downS == Wall) {
		ko_flag = true
	}
	if prisonersCount == 1 && ko_flag {
		pos.KoStone = move.Stone
		if pos.GetStone(up) == Empty {
			pos.Ko = up
		} else if pos.GetStone(down) == Empty {
			pos.Ko = down
		} else if pos.GetStone(left) == Empty {
			pos.Ko = left
		} else if pos.GetStone(right) == Empty {
			pos.Ko = right
		}
	} else {
		pos.Ko, pos.KoStone = Outbound, Empty
	}
}
