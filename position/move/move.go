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

package move

import (
	"github.com/foozea/isana/board/stone"
	"github.com/foozea/isana/board/vertex"
)

type Move struct {
	Stone  stone.Stone
	Vertex vertex.Vertex
	///
	Games     int32
	Rate      float64
	RaveGames int32
	RaveRate  float64
	UCB       float64
}

var PassMove Move = Move{stone.Empty, vertex.Outbound, 0, 0.0, 0, 0, 0}

func CreateMove(s stone.Stone, v vertex.Vertex) *Move {
	return &Move{s, v, 0, 0.0, 0, 0, 0}
}

func (m *Move) String() string {
	return m.Vertex.String()
}
