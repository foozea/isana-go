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

package hashing

import (
	. "github.com/foozea/isana/board/stone"
)

type Category int

const (
	Eye Category = 1
)

var Patterns = make(map[uint64]Category, 0)

func init() {
	hash := func(s Stone, is ...int) uint64 {
		code := uint64(0x0)
		for _, v := range is {
			code ^= DeltaHash[v<<2|int(s)]
		}
		return code
	}

	// Eye patterns
	// Black
	Patterns[hash(Black, 1, 2, 3, 4, 6, 7, 8, 9)] = Eye
	Patterns[hash(Black, 2, 4, 6, 8)] = Eye
	Patterns[hash(Black, 2, 3, 4, 6, 7, 8, 9)^hash(White, 1)] = Eye
	Patterns[hash(Black, 1, 2, 4, 6, 7, 8, 9)^hash(White, 3)] = Eye
	Patterns[hash(Black, 1, 2, 3, 4, 6, 8, 9)^hash(White, 7)] = Eye
	Patterns[hash(Black, 1, 2, 3, 4, 6, 7, 8)^hash(White, 9)] = Eye
	Patterns[hash(Black, 2, 3, 4, 6, 8, 9)^hash(White, 1)] = Eye
	Patterns[hash(Black, 1, 2, 4, 6, 7, 8)^hash(White, 3)] = Eye
	Patterns[hash(Black, 2, 3, 4, 6, 8, 9)^hash(White, 7)] = Eye
	Patterns[hash(Black, 1, 2, 4, 6, 7, 8)^hash(White, 9)] = Eye
	Patterns[hash(Black, 2, 3, 4, 6, 8)^hash(White, 1)] = Eye
	Patterns[hash(Black, 2, 4, 6, 7, 8)^hash(White, 3)] = Eye
	Patterns[hash(Black, 2, 3, 4, 6, 8)^hash(White, 7)] = Eye
	Patterns[hash(Black, 2, 4, 6, 7, 8)^hash(White, 9)] = Eye
	Patterns[hash(Black, 2, 4, 6, 8, 9)^hash(White, 1)] = Eye
	Patterns[hash(Black, 1, 2, 4, 6, 8)^hash(White, 3)] = Eye
	Patterns[hash(Black, 2, 4, 6, 8, 9)^hash(White, 7)] = Eye
	Patterns[hash(Black, 1, 2, 4, 6, 8)^hash(White, 9)] = Eye
	Patterns[hash(Black, 2, 4, 6, 8)^hash(White, 1)] = Eye
	Patterns[hash(Black, 2, 4, 6, 8)^hash(White, 3)] = Eye
	Patterns[hash(Black, 2, 4, 6, 8)^hash(White, 7)] = Eye
	Patterns[hash(Black, 2, 4, 6, 8)^hash(White, 9)] = Eye
	//White
	Patterns[hash(White, 1, 2, 3, 4, 6, 7, 8, 9)] = Eye
	Patterns[hash(White, 2, 4, 6, 8)] = Eye
	Patterns[hash(White, 2, 3, 4, 6, 7, 8, 9)^hash(Black, 1)] = Eye
	Patterns[hash(White, 1, 2, 4, 6, 7, 8, 9)^hash(Black, 3)] = Eye
	Patterns[hash(White, 1, 2, 3, 4, 6, 8, 9)^hash(Black, 7)] = Eye
	Patterns[hash(White, 1, 2, 3, 4, 6, 7, 8)^hash(Black, 9)] = Eye
	Patterns[hash(White, 2, 3, 4, 6, 8, 9)^hash(Black, 1)] = Eye
	Patterns[hash(White, 1, 2, 4, 6, 7, 8)^hash(Black, 3)] = Eye
	Patterns[hash(White, 2, 3, 4, 6, 8, 9)^hash(Black, 7)] = Eye
	Patterns[hash(White, 1, 2, 4, 6, 7, 8)^hash(Black, 9)] = Eye
	Patterns[hash(White, 2, 3, 4, 6, 8)^hash(Black, 1)] = Eye
	Patterns[hash(White, 2, 4, 6, 7, 8)^hash(Black, 3)] = Eye
	Patterns[hash(White, 2, 3, 4, 6, 8)^hash(Black, 7)] = Eye
	Patterns[hash(White, 2, 4, 6, 7, 8)^hash(Black, 9)] = Eye
	Patterns[hash(White, 2, 4, 6, 8, 9)^hash(Black, 1)] = Eye
	Patterns[hash(White, 1, 2, 4, 6, 8)^hash(Black, 3)] = Eye
	Patterns[hash(White, 2, 4, 6, 8, 9)^hash(Black, 7)] = Eye
	Patterns[hash(White, 1, 2, 4, 6, 8)^hash(Black, 9)] = Eye
	Patterns[hash(White, 2, 4, 6, 8)^hash(Black, 1)] = Eye
	Patterns[hash(White, 2, 4, 6, 8)^hash(Black, 3)] = Eye
	Patterns[hash(White, 2, 4, 6, 8)^hash(Black, 7)] = Eye
	Patterns[hash(White, 2, 4, 6, 8)^hash(Black, 9)] = Eye
}
