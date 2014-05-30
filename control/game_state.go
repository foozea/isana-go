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

package control

import (
	. "code.isana.io/isana/board/size"
	. "code.isana.io/isana/board/stone"
	. "code.isana.io/isana/board/vertex"
	. "code.isana.io/isana/position"
	. "code.isana.io/isana/position/move"
)

type GameState struct {
	Size    BoardSize
	Current *Position
	Turn    Stone
	Komi    float64
	Times   TimeSettings
}

type TimeSettings struct {
	MainTime      int
	ByoYomiTime   int
	ByoYomiStones int
}

func CreateDefaultGameState() GameState {
	komi := 0.0
	timeset := TimeSettings{60, 600, 25}
	return GameState{B9x9, nil, Empty, komi, timeset}
}

func (s *GameState) GetCurrentPosition() *Position {
	if s.Turn == Empty {
		pos := CreatePosition(s.Size)
		return &pos
	}
	return s.Current
}

func (s *GameState) CurrentStoneIs() Stone {
	return s.Turn
}

func (s *GameState) MakeMove(move *Move) bool {
	pos := s.GetCurrentPosition()
	next, ok := pos.PseudoMove(move)
	if !ok {
		return false
	}
	next.FixMove(move)
	s.Current = next
	s.Turn = move.Stone
	return true
}

func (s *GameState) Pass() {
	current := s.GetCurrentPosition()
	next := CopyPosition(current)
	next.Ko, next.KoStone = Outbound, Empty
	s.Current = &next
	s.Turn = s.Turn.Opposite()
}

func (s *GameState) ClearHistory() {
	s.Current = nil
	s.Turn = Empty
}
