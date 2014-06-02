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

package bitboard

import (
	. "github.com/foozea/isana/board/size"
)

// Type: bitboard
// this is the backend structure of Go-board.
type Bitboard [6]uint64

// Bitmasks that restricts the range of bit-shift.
// these are used in:
//   - Left, Right, Up, Down
var (
	bitmask9l = Bitboard{
		0xbfdfeff7fbfdfeff, 0xff7f, 0x0, 0x0, 0x0, 0x0}
	bitmask9r = Bitboard{
		0x7fbfdfeff7fbfdfe, 0x1feff, 0x0, 0x0, 0x0, 0x0}
	bitmask9u = Bitboard{
		0xffffffffffffffff, 0x1ffff, 0x0, 0x0, 0x0, 0x0}

	bitmask11l = Bitboard{
		0xffbff7feffdffbff, 0xeffdffbff7feffd, 0x0, 0x0, 0x0, 0x0}
	bitmask11r = Bitboard{
		0xff7feffdffbff7fe, 0xfffbff7feffdffb, 0x0, 0x0, 0x0, 0x0}
	bitmask11u = Bitboard{
		0xffffffffffffffff, 0x1ffffffffffffff, 0x0, 0x0, 0x0, 0x0}

	bitmask13l = Bitboard{
		0xfff7ffbffdffefff, 0xffefff7ffbffdffe, 0xfff7ffbffd, 0x0, 0x0, 0x0}
	bitmask13r = Bitboard{
		0xffefff7ffbffdffe, 0xffdffefff7ffbffd, 0xfffefff7ffb, 0x0, 0x0, 0x0}
	bitmask13u = Bitboard{
		0xffffffffffffffff, 0xffffffffffffffff, 0x1ffffffffff, 0x0, 0x0, 0x0}

	bitmask15l = Bitboard{
		0xf7ffefffdfffbfff, 0xff7ffefffdfffbff,
		0xfff7ffefffdfffbf, 0xfffdfffb, 0x0, 0x0}
	bitmask15r = Bitboard{
		0xefffdfffbfff7ffe, 0xfefffdfffbfff7ff,
		0xffefffdfffbfff7f, 0x1fffbfff7, 0x0, 0x0}
	bitmask15u = Bitboard{
		0xffffffffffffffff, 0xffffffffffffffff, 0xffffffffffffffff,
		0x1ffffffff, 0x0, 0x0}

	bitmask19l = Bitboard{
		0xfeffffdffffbffff, 0xfffdffffbffff7ff, 0xdffffbffff7fffef,
		0xffbffff7fffeffff, 0xffff7fffeffffdff, 0xeffffdffffb}
	bitmask19r = Bitboard{
		0xfdffffbffff7fffe, 0xfffbffff7fffefff, 0xbffff7fffeffffdf,
		0xff7fffeffffdffff, 0xfffeffffdffffbff, 0xfffffbffff7}
	bitmask19u = Bitboard{
		0xffffffffffffffff, 0xffffffffffffffff, 0xffffffffffffffff,
		0xffffffffffffffff, 0xffffffffffffffff, 0x1ffffffffff}
)

// Gets a bit of the index.
func (b *Bitboard) GetBit(index int) int {
	sidx := index / 64 // slice index
	bidx := index % 64 // bit index
	return int((*b)[sidx]>>uint(bidx)) & 0x1
}

// Set a bit of the index.
func (b *Bitboard) SetBit(index int) {
	sidx := index / 64 // slice index
	bidx := index % 64 // bit index
	(*b)[sidx] |= uint64(0x1 << uint(bidx))
}

// Clears a bit of the index.
func (b *Bitboard) ClearBit(index int) {
	sidx := index / 64 // slice index
	bidx := index % 64 // bit index
	(*b)[sidx] &= ^(0x1 << uint(bidx))
}

// Counts the bit number that set on the board.
func (b *Bitboard) CountBit() int {
	ret := 0
	for _, v := range *b {
		n := v
		for n > 0 {
			n &= n - 1
			ret++
		}
	}
	return ret
}

// Bit-wize operation: `AND`
func And(lhs Bitboard, rhs ...Bitboard) Bitboard {
	calc := Bitboard{}
	for i, v := range lhs {
		calc[i] = v
		for _, b := range rhs {
			calc[i] &= b[i]
		}
	}
	return calc
}

// Bit-wize operation: `OR`
func Or(lhs Bitboard, rhs ...Bitboard) Bitboard {
	calc := Bitboard{}
	for i, v := range lhs {
		calc[i] = v
		for _, b := range rhs {
			calc[i] |= b[i]
		}
	}
	return calc
}

// Bit-wize operation: `XOR`
func Xor(lhs Bitboard, rhs ...Bitboard) Bitboard {
	calc := Bitboard{}
	for i, v := range lhs {
		calc[i] = v
		for _, b := range rhs {
			calc[i] ^= b[i]
		}
	}
	return calc
}

// Bit-wize operation: `NOT`
func Not(bits Bitboard) Bitboard {
	calc := Bitboard{}
	for i, v := range bits {
		calc[i] = ^v
	}
	return calc
}

// One-bit left shift. the shift size is constant but
// the result is masked with the bits that depends on the board-size.
func Left(bits Bitboard, size BoardSize) Bitboard {
	_1, _2, _3, _4, _5 := bits[1]<<63, bits[2]<<63, bits[3]<<63, bits[4]<<63, bits[5]>>63
	calc := Bitboard{_1, _2, _3, _4, _5, 0x0}
	for i, v := range bits {
		calc[i] |= v >> 1
		switch size {
		case B9x9:
			calc[i] &= bitmask9l[i]
		case B11x11:
			calc[i] &= bitmask11l[i]
		case B13x13:
			calc[i] &= bitmask13l[i]
		case B15x15:
			calc[i] &= bitmask15l[i]
		case B19x19:
			calc[i] &= bitmask19l[i]
		}
	}
	return calc
}

// One-bit right shift. the shift size is constant but
// the result is masked with the bits that depends on the board-size.
func Right(bits Bitboard, size BoardSize) Bitboard {
	_0, _1, _2, _3, _4 := bits[0]>>63, bits[1]>>63, bits[2]>>63, bits[3]>>63, bits[4]>>63
	calc := Bitboard{0x0, _0, _1, _2, _3, _4}
	for i, v := range bits {
		calc[i] |= v << 1
		switch size {
		case B9x9:
			calc[i] &= bitmask9r[i]
		case B11x11:
			calc[i] &= bitmask11r[i]
		case B13x13:
			calc[i] &= bitmask13r[i]
		case B15x15:
			calc[i] &= bitmask15r[i]
		case B19x19:
			calc[i] &= bitmask19r[i]
		}
	}
	return calc
}

// One-bit up shift (<line-size> bit left shift).
// the result is masked with the bits that depends on the board-size.
func Up(bits Bitboard, size BoardSize) Bitboard {
	m := 64 - uint(size)
	_0, _1, _2, _3, _4 := bits[0]>>m, bits[1]>>m, bits[2]>>m, bits[3]>>m, bits[4]>>m
	calc := Bitboard{0x0, _0, _1, _2, _3, _4}
	for i, v := range bits {
		calc[i] |= v << uint(size)
		switch size {
		case B9x9:
			calc[i] &= bitmask9u[i]
		case B11x11:
			calc[i] &= bitmask11u[i]
		case B13x13:
			calc[i] &= bitmask13u[i]
		case B15x15:
			calc[i] &= bitmask15u[i]
		case B19x19:
			calc[i] &= bitmask19u[i]
		}
	}
	return calc
}

// One-bit up shift (<line-size> bit right shift).
// the result is masked with the bits that depends on the board-size.
func Down(bits Bitboard, size BoardSize) Bitboard {
	m := 64 - uint(size)
	_1, _2, _3, _4, _5 := bits[1]<<m, bits[2]<<m, bits[3]<<m, bits[4]<<m, bits[5]>>m
	calc := Bitboard{_1, _2, _3, _4, _5, 0x0}
	for i, v := range bits {
		calc[i] |= v >> uint(size)
	}
	return calc
}
