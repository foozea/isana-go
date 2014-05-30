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

import (
	"testing"
)

func TestCapacity(t *testing.T) {
	var actual, expected int
	const msg string = "Capacity / couldn't get capacity. expected : %v, but %v"
	actual = B9x9.Capacity()
	expected = 81
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = B11x11.Capacity()
	expected = 121
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = B13x13.Capacity()
	expected = 169
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = B15x15.Capacity()
	expected = 225
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = B19x19.Capacity()
	expected = 361
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = BoardSize(999).Capacity()
	expected = 0
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}
