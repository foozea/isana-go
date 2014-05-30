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

package protocol

import (
	. "github.com/foozea/isana/board/size"
	. "github.com/foozea/isana/board/stone"
	. "github.com/foozea/isana/board/vertex"
	. "github.com/foozea/isana/position"
	. "github.com/foozea/isana/position/move"

	"fmt"
	"os"
	"strconv"
)

type Args []string

func (a *Args) Clear() {
	*a = make(Args, 2)
}

type Handler func(Args)

func protocol_version(args Args) {
	fmt.Printf("= %v\n\n", PROTOCOL_VERSION)
}

func name(args Args) {
	fmt.Printf("= %v\n\n", Engine.Name)
}

func version(args Args) {
	fmt.Printf("= %v\n\n", Engine.Version)
}

func known_command(args Args) {
	fmt.Print("= ")
	if len(args) == 0 {
		fmt.Println("false\n")
	} else {
		fmt.Printf("%v\n\n", Dispatcher.HasHandler(args[0]))
	}
}

func list_commands(args Args) {
	fmt.Print("= ")
	for k, _ := range Dispatcher {
		fmt.Println(k)
	}
	fmt.Printf("\n")
}

func quit(args Args) {
	os.Exit(0)
}

func boardsize(args Args) {
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
		GameController.Size = B9x9
	case 11:
		GameController.Size = B11x11
	case 13:
		GameController.Size = B13x13
	case 15:
		GameController.Size = B15x15
	case 19:
		GameController.Size = B19x19
	default:
		fmt.Println("? unacceptable size\n")
		return
	}
	CreateHash(GameController.Size)
	clear_board(args)
}

func clear_board(args Args) {
	GameController.ClearHistory()
	fmt.Println("=\n")
}

func komi(args Args) {
	if len(args) == 0 {
		fmt.Println("? komi must be a float\n")
		return
	}
	v, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Println("? komi must be a float\n")
		return
	}
	GameController.Komi = v
	Engine.Komi = v
	fmt.Println("=\n")
}

func play(args Args) {
	if len(args) < 2 {
		fmt.Println("? invalid parameter(s)\n")
		return
	}
	stone := StringToStone(args[0])
	if stone == Wall || stone == Empty {
		fmt.Println("? invalid parameter(s)\n")
		return
	}
	point := StringToVertex(args[1], GameController.Size)
	if point == Outbound { //pass
		GameController.Pass()
		fmt.Println("=\n")
		return
	}
	mv := CreateMove(stone, point)
	ok := GameController.MakeMove(mv)
	if !ok {
		fmt.Println("? illegal move\n")
		return
	}
	fmt.Println("=\n")
}

func genmove(args Args) {
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
	pos := GameController.GetCurrentPosition()
	mv := Engine.Ponder(pos, stone)
	ret := GameController.MakeMove(&mv)
	if !ret {
		fmt.Printf("= PASS\n\n")
	}
	fmt.Printf("= %v\n\n", mv.String())
}

func showboard(args Args) {
	fmt.Println("=\n")
	pos := GameController.GetCurrentPosition()
	pos.Dump()
	pos.GoStringDump()
	fmt.Printf("Black (X) : %v stones\n", pos.BlackPrison)
	fmt.Printf("White (O) : %v stones\n\n", pos.WhitePrison)
}
