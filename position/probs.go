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

	. "math"
	. "math/rand"
	. "time"
)

func init() {
	Seed(Now().UTC().UnixNano())
}

func (pos *Position) CreateProbs() {
	pos.TotalProbs = 0
	subtotal := 0
	bits := Or(pos.blacks, pos.whites)
	for i := 0; i < pos.Size.Capacity(); i++ {
		if bits.GetBit(i) == 1 {
			pos.ProbDencities[i] = 0
		} else {
			n := Intn(MaxInt8)
			pos.ProbDencities[i] = n
			pos.TotalProbs += n
			subtotal += n
		}
		if i%int(pos.Size) == 0 {
			pos.SubTotalProbs[i/int(pos.Size)] = subtotal
			subtotal = 0
		}
	}
}

func (pos *Position) SearchProbIndex(value int) int {
	var selected, v int
	for selected, v = range pos.SubTotalProbs {
		if v >= value {
			break
		}
		value = value - v
	}
	for i := selected * int(pos.Size); i < len(pos.ProbDencities); i++ {
		v := pos.ProbDencities[i]
		if v == 0 {
			continue
		}
		value -= v
		if value <= 0 {
			return i
		}
	}
	return pos.ProbDencities[pos.Size.Capacity()-1]
}

func (pos *Position) UpdateProbs(key int, value int) {
	current := pos.ProbDencities[key]
	delta := value - current
	pos.TotalProbs += delta
	pos.SubTotalProbs[key/int(pos.Size)] -= delta
	pos.ProbDencities[key] = value
}
