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

package main

import (
	"flag"
	"runtime"

	. "github.com/foozea/isana/engine"
	"github.com/foozea/isana/protocol/gtp"
)

// Defines engine name and version.
const (
	name    string = "Isana"
	version string = "0.1"
)

var (
	processNumber int
	parallelRoots int
	trialNumber   int
)

func init() {
	// parse flags
	flag.IntVar(&processNumber, "processes", runtime.NumCPU(), "process number")
	flag.IntVar(&processNumber, "p", runtime.NumCPU(), "process number")
	flag.IntVar(&parallelRoots, "roots", 3, "root parallelize number")
	flag.IntVar(&parallelRoots, "r", 3, "root parallelize number")
	flag.IntVar(&trialNumber, "trials", 3000, "uct trial number")
	flag.IntVar(&trialNumber, "t", 3000, "uct trial number")
	flag.Parse()

	runtime.GOMAXPROCS(processNumber)

	Engine.Name = name
	Engine.Version = version
	Engine.Trials = trialNumber / parallelRoots
	Engine.Roots = parallelRoots
}

func main() {
	gtp.Start()
}
