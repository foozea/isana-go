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

type Args []string
type Handler func(Args)
type Command string
type CommandMap map[string]Handler

var (
	Dispatcher      CommandMap
	ArgsForHandlers Args
)

func init() {
	Dispatcher = make(map[string]Handler, 0)
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
