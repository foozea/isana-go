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

package vertex

import (
	"strconv"
	"strings"

	. "github.com/foozea/isana/board/size"
)

type Vertex struct {
	Index int
	Size  BoardSize
}

var Outbound = Vertex{-1, BoardSize(0)}

func (vx Vertex) IsValid() bool {
	return vx.Size.Capacity() > vx.Index && vx.Index >= 0
}

func (vx Vertex) Up() Vertex {
	v := vx.Index + int(vx.Size)
	return Vertex{v, vx.Size}
}

func (vx Vertex) Down() Vertex {
	v := vx.Index - int(vx.Size)
	return Vertex{v, vx.Size}
}

func (vx Vertex) Left() Vertex {
	if vx == Outbound || vx.Index%int(vx.Size) == 0 {
		return Outbound
	}
	v := vx.Index - 1
	return Vertex{v, vx.Size}
}

func (vx Vertex) Right() Vertex {
	if vx == Outbound || vx.Index%int(vx.Size) == int(vx.Size)-1 {
		return Outbound
	}
	v := vx.Index + 1
	return Vertex{v, vx.Size}
}

func (vx Vertex) String() string {
	if int(vx.Size) == 0 {
		return "PASS"
	}
	file := vx.Index % int(vx.Size)
	rank := vx.Index / int(vx.Size)
	if vx != Outbound {
		f := string('A' + file)
		if f == "I" {
			f = "J"
		}
		r := strconv.Itoa(rank + 1)
		return f + r
	} else {
		return "PASS"
	}
}

func StringToVertex(str string, size BoardSize) Vertex {
	str = strings.ToUpper(str)
	if len(str) < 2 || str == "PASS" {
		return Outbound
	}
	file := int(([]rune(str[:1]))[0] - 'A')
	if file > 8 {
		file -= 1
	}
	rank, err := strconv.Atoi(str[1:len(str)])
	if err == nil {
		vx := Vertex{(rank-1)*int(size) + file, size}
		if vx.IsValid() {
			return Vertex{(rank-1)*int(size) + file, size}
		}
	}
	return Outbound
}
