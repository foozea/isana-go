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

package misc

import (
	"log"
	"time"
)

func Trace(s string) (string, time.Time) {
	log.Println("START:", s)
	return s, time.Now()
}

func Un(s string, startTime time.Time) {
	elapsed := time.Since(startTime)
	log.Printf("TRACE END: %s, elapsed time: %f secs\n", s, elapsed.Seconds())
}
