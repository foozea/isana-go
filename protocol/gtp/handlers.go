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

package gtp

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	. "github.com/foozea/isana/board/size"
	. "github.com/foozea/isana/board/stone"
	. "github.com/foozea/isana/board/vertex"
	. "github.com/foozea/isana/position"
	. "github.com/foozea/isana/position/move"
	"github.com/foozea/isana/protocol"
)

var (
	Roots int
)

func init() {
	Roots = 1
}

func protocol_version(args protocol.Args) {
	fmt.Printf("= %v\n\n", PROTOCOL_VERSION)
}

func name(args protocol.Args) {
	fmt.Printf("= %v\n\n", protocol.Engine.Name)
}

func version(args protocol.Args) {
	fmt.Printf("= %v\n\n", protocol.Engine.Version)
}

func known_command(args protocol.Args) {
	fmt.Print("= ")
	if len(args) == 0 {
		fmt.Println("false\n")
	} else {
		fmt.Printf("%v\n\n", protocol.Dispatcher.HasHandler(args[0]))
	}
}

func list_commands(args protocol.Args) {
	fmt.Print("= ")
	for k, _ := range protocol.Dispatcher {
		fmt.Println(k)
	}
	fmt.Printf("\n")
}

func quit(args protocol.Args) {
	os.Exit(0)
}

func boardsize(args protocol.Args) {
	if len(args) == 0 {
		fmt.Println("? boardsize must be an integer\n")
		return
	}
	v, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("? boardsize must be an integer\n")
		return
	}
	switch v {
	case 9:
		protocol.GameController.Size = B9x9
	case 11:
		protocol.GameController.Size = B11x11
	case 13:
		protocol.GameController.Size = B13x13
	case 15:
		protocol.GameController.Size = B15x15
	case 19:
		protocol.GameController.Size = B19x19
	default:
		fmt.Println("? unacceptable size\n")
		return
	}
	clear_board(args)
}

func clear_board(args protocol.Args) {
	protocol.GameController.ClearHistory()
	fmt.Println("=\n")
}

func komi(args protocol.Args) {
	if len(args) == 0 {
		fmt.Println("? komi must be a float\n")
		return
	}
	v, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Println("? komi must be a float\n")
		return
	}
	protocol.GameController.Komi = v
	protocol.Engine.Komi = v
	fmt.Println("=\n")
}

func play(args protocol.Args) {
	if len(args) < 2 {
		fmt.Println("? invalid parameter(s)\n")
		return
	}
	stone := StringToStone(args[0])
	if stone == Wall || stone == Empty {
		fmt.Println("? invalid parameter(s)\n")
		return
	}
	point := StringToVertex(args[1], protocol.GameController.Size)
	if point == Outbound { //pass
		protocol.GameController.Pass()
		fmt.Println("=\n")
		return
	}
	mv := CreateMove(stone, point)
	ok := protocol.GameController.MakeMove(mv)
	if !ok {
		fmt.Println("? illegal move\n")
		return
	}
	fmt.Println("=\n")
}

func genmove(args protocol.Args) {
	if len(args) == 0 {
		fmt.Println("? invalid parameter(s)\n")
		return
	}
	stone := StringToStone(args[0])
	if stone == Wall {
		fmt.Println("? invalid parameter(s)\n")
		return
	}
	if stone == Empty { // resign
		fmt.Println("= RESIGN\n")
		return
	}
	pos := protocol.GameController.GetCurrentPosition()

	// Root Parallelize
	mvs := make([]Move, Roots)
	var wg sync.WaitGroup
	for i, _ := range mvs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			copied := CopyPosition(pos)
			mvs[i] = protocol.Engine.Think(&copied, stone)
		}()
	}
	wg.Wait()

	selected := mvs[0]
	for _, v := range mvs {
		if selected.Rate < v.Rate {
			selected = v
		}
	}

	ret := protocol.GameController.MakeMove(&selected)
	if !ret {
		fmt.Printf("= PASS\n\n")
		return
	}
	fmt.Printf("= %v\n\n", selected.String())
}

func showboard(args protocol.Args) {
	fmt.Println("=\n")
	pos := protocol.GameController.GetCurrentPosition()
	pos.Dump()
	pos.GoStringDump()
	fmt.Printf("Black (X) : %v stones\n", pos.BlackPrison)
	fmt.Printf("White (O) : %v stones\n\n", pos.WhitePrison)
}
