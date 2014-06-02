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
	. "github.com/foozea/isana/board/stone"
	. "github.com/foozea/isana/hashing"
	. "github.com/foozea/isana/position/move"
)

/// Validations for PseudoMove

// Determines if the move is valid or not.
func (pos *Position) isLegalMove(move *Move) bool {
	// 1. if it is not empty, returns false
	if pos.GetStone(move.Vertex) != Empty {
		return false
	}
	// 2. if it is Ko threat or the vertex is not empty, returns false
	if move.Stone == pos.KoStone.Opposite() && move.Vertex == pos.Ko {
		return false
	}
	return true
}

// Determines if the move fills own eye.
func (pos *Position) isFillEyeMove(move *Move) bool {
	hash := pos.SquaredHash3(move.Vertex)
	category := Patterns[hash]
	if category != 0 && category == Eye {
		return true
	}
	return false
}

// Determines if the move is suicide or not.
// <border> is the border number to evaluate whether the move is suicide.
func (pos *Position) isSuicideMove(move *Move, border int) bool {
	stone := move.Stone
	vx := move.Vertex
	// 1. check the tempolary position and if the stone is not dead,
	//    it is not suicide move.
	if pos.CountLiberty(pos.GetString(vx)) > border {
		return false
	}
	// 2. if the move can take opponent stone(s), not suicide.
	up, down, left, right := vx.Up(), vx.Down(), vx.Left(), vx.Right()
	opp := stone.Opposite()
	if (pos.GetStone(left) == opp && pos.CountLiberty(pos.GetString(left)) == 0) ||
		(pos.GetStone(right) == opp && pos.CountLiberty(pos.GetString(right)) == 0) ||
		(pos.GetStone(up) == opp && pos.CountLiberty(pos.GetString(up)) == 0) ||
		(pos.GetStone(down) == opp && pos.CountLiberty(pos.GetString(down)) == 0) {
		return false
	}
	// 3. else, suicide move.
	return true
}
