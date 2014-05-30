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
	. "github.com/foozea/isana/control"
	. "github.com/foozea/isana/engine"
)

type Command string
type CommandMap map[string]Handler

const (
	Protocol_version Command = "protocol_version"
	Name             Command = "name"
	Version          Command = "version"
	Known_command    Command = "known_command"
	List_commands    Command = "list_commands"
	Quit             Command = "quit"
	Boardsize        Command = "boardsize"
	Clear_board      Command = "clear_board"
	Komi             Command = "komi"
	Play             Command = "play"
	Genmove          Command = "genmove"
	Showboard        Command = "showboard"
)

const (
	PROTOCOL_VERSION int = 2
	COMMANDS_COUNT   int = 10
)

var (
	Engine          Isana
	Dispatcher      CommandMap
	ArgsForHandlers Args
	GameController  GameState
)

func init() {
	Engine = CreateEngine("", "")

	Dispatcher = make(map[string]Handler, COMMANDS_COUNT)
	ArgsForHandlers = make(Args, 5)
	GameController = CreateDefaultGameState()

	// register handlers
	Dispatcher.AddHandler(Protocol_version, protocol_version)
	Dispatcher.AddHandler(Name, name)
	Dispatcher.AddHandler(Version, version)
	Dispatcher.AddHandler(Known_command, known_command)
	Dispatcher.AddHandler(List_commands, list_commands)
	Dispatcher.AddHandler(Boardsize, boardsize)
	Dispatcher.AddHandler(Clear_board, clear_board)
	Dispatcher.AddHandler(Komi, komi)
	Dispatcher.AddHandler(Play, play)
	Dispatcher.AddHandler(Genmove, genmove)
	Dispatcher.AddHandler(Quit, quit)
	Dispatcher.AddHandler(Showboard, showboard)
}

func (m CommandMap) AddHandler(key Command, handler func(Args)) {
	m[key.String()] = Handler(handler)
}

func (m CommandMap) HasHandler(key string) bool {
	if m[key] != nil {
		return true
	}
	return false
}

func (m CommandMap) CallHandler(key string) {
	h := m[key]
	Handler(h)(ArgsForHandlers)
}

func (c Command) String() string {
	return string(c)
}
