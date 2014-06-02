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

package size

type BoardSize uint

// Definitions of the board-sizes.
const (
	B9x9   BoardSize = 9
	B11x11 BoardSize = 11
	B13x13 BoardSize = 13
	B15x15 BoardSize = 15
	B19x19 BoardSize = 19
)

// Returns the vertex count that can be set a stone.
func (size BoardSize) Capacity() int {
	switch size {
	case B9x9:
		return 81
	case B11x11:
		return 121
	case B13x13:
		return 169
	case B15x15:
		return 225
	case B19x19:
		return 361
	}
	return 0
}
