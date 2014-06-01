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
	"testing"

	. "github.com/foozea/isana/board/size"
)

var (
	board Bitboard
)

func TestSetBit(t *testing.T) {
	var actual, expected uint64
	const msg string = "SetBit / failed to set bit. expected : %v, but %v"
	board.SetBit(3)
	actual = board[0]
	expected = 0x8
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	board.SetBit(80)
	actual = board[1]
	expected = uint64(65536)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}

func TestGetBit(t *testing.T) {
	var actual, expected int
	const msg string = "GetBit / failed to get bit. expected : %v, but %v"
	board.SetBit(3)
	board.SetBit(80)
	board.SetBit(130)
	board.SetBit(360)
	actual = board.GetBit(3)
	expected = 1
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = board.GetBit(80)
	expected = 1
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = board.GetBit(130)
	expected = 1
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = board.GetBit(360)
	expected = 1
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = board.GetBit(359)
	expected = 0
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}

func TestClearBit(t *testing.T) {
	var actual, expected int
	const msg string = "ClearBit / failed to clear bit. expected : %v, but %v"
	board.SetBit(360)
	actual = board.GetBit(360)
	expected = 1
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	board.ClearBit(360)
	actual = board.GetBit(360)
	expected = 0
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}

func TestCountBit(t *testing.T) {
	var actual, expected int
	const msg string = "CountBit / failed to get bit. expected : %v, but %v"
	counted := Bitboard{}
	counted.SetBit(1)
	counted.SetBit(3)
	counted.SetBit(5)
	counted.SetBit(7)
	counted.SetBit(9)
	actual = counted.CountBit()
	expected = 5
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}

func TestAnd(t *testing.T) {
	var actual, expected Bitboard
	const msg string = "And / Bit-wise operation `And` fialed. expected : %v, but %v"
	lhs := Bitboard{}
	rhs := Bitboard{}
	lhs.SetBit(1)
	rhs.SetBit(10)
	actual = And(lhs, rhs)
	expected = Bitboard{}
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}

func TestOr(t *testing.T) {
	var actual, expected Bitboard
	const msg string = "Or / Bit-wise operation `Or` fialed. expected : %v, but %v"
	lhs := Bitboard{}
	rhs := Bitboard{}
	lhs.SetBit(1)
	rhs.SetBit(10)
	actual = Or(lhs, rhs)
	expected = Bitboard{uint64(1026), 0x0, 0x0, 0x0, 0x0, 0x0}
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}

func TestXor(t *testing.T) {
	var actual, expected Bitboard
	const msg string = "Xor / Bit-wise operation `Xor` fialed. expected : %v, but %v"
	lhs := Bitboard{}
	rhs := Bitboard{}
	lhs.SetBit(1)
	rhs.SetBit(1)
	rhs.SetBit(10)
	actual = Xor(lhs, rhs)
	expected = Bitboard{uint64(1024), 0x0, 0x0, 0x0, 0x0, 0x0}
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}

func TestNot(t *testing.T) {
	var actual, expected Bitboard
	const msg string = "Not / Bit-wise operation `Not` fialed. expected : %v, but %v"
	lhs := Bitboard{}
	lhs.SetBit(1)
	actual = Not(lhs)
	expected = Bitboard{18446744073709551613, 18446744073709551615,
		18446744073709551615, 18446744073709551615, 18446744073709551615, 18446744073709551615}
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}

func TestRight(t *testing.T) {
	var actual, expected Bitboard
	const msg string = "Right / 1 bit right shift is not works. expected : %v, but %v"
	b := Bitboard{}
	b.SetBit(63)
	actual = Right(b, B9x9)
	expected = Bitboard{}
	expected.SetBit(64)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = Right(b, B11x11)
	expected = Bitboard{}
	expected.SetBit(64)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = Right(b, B13x13)
	expected = Bitboard{}
	expected.SetBit(64)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = Right(b, B15x15)
	expected = Bitboard{}
	expected.SetBit(64)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = Right(b, B19x19)
	expected = Bitboard{}
	expected.SetBit(64)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}

func TestLeft(t *testing.T) {
	var actual, expected Bitboard
	const msg string = "Left / 1 bit left shift is not works. expected : %v, but %v"
	b := Bitboard{}
	b.SetBit(76)
	actual = Left(b, B9x9)
	expected = Bitboard{}
	expected.SetBit(75)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = Left(b, B11x11)
	expected = Bitboard{}
	expected.SetBit(75)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = Left(b, B13x13)
	expected = Bitboard{}
	expected.SetBit(75)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = Left(b, B15x15)
	expected = Bitboard{}
	expected.SetBit(75)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	b.ClearBit(76)
	b.SetBit(78)
	actual = Left(b, B19x19)
	expected = Bitboard{}
	expected.SetBit(77)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}

func TestUp(t *testing.T) {
	var actual, expected Bitboard
	const msg string = "Up / 1 bit up shift is not works. expected : %v, but %v"
	b := Bitboard{}
	b.SetBit(62)
	actual = Up(b, B9x9)
	expected = Bitboard{}
	expected.SetBit(71)
	expected = And(expected, bitmask9r)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = Up(b, B11x11)
	expected = Bitboard{}
	expected.SetBit(73)
	expected = And(expected, bitmask11r)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = Up(b, B13x13)
	expected = Bitboard{}
	expected.SetBit(75)
	expected = And(expected, bitmask13r)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = Up(b, B15x15)
	expected = Bitboard{}
	expected.SetBit(77)
	expected = And(expected, bitmask15r)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = Up(b, B19x19)
	expected = Bitboard{}
	expected.SetBit(81)
	expected = And(expected, bitmask19r)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}

func TestDown(t *testing.T) {
	var actual, expected Bitboard
	const msg string = "Down / 1 bit down shift is not works. expected : %v, but %v"
	b := Bitboard{}
	b.SetBit(63)
	actual = Down(b, B9x9)
	expected = Bitboard{}
	expected.SetBit(54)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = Down(b, B11x11)
	expected = Bitboard{}
	expected.SetBit(52)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = Down(b, B13x13)
	expected = Bitboard{}
	expected.SetBit(50)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = Down(b, B15x15)
	expected = Bitboard{}
	expected.SetBit(48)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = Down(b, B19x19)
	expected = Bitboard{}
	expected.SetBit(44)
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}
