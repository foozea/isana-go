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

package engine

import (
	"log"

	. "math"
	. "math/rand"
	. "sync"
	. "time"

	. "github.com/foozea/isana/board/stone"
	. "github.com/foozea/isana/board/vertex"
	. "github.com/foozea/isana/misc"
	. "github.com/foozea/isana/position"
	. "github.com/foozea/isana/position/move"
)

var mutex = &Mutex{}

func init() {
	Seed(Now().UTC().UnixNano())
}

type Isana struct {
	Name            string
	Version         string
	Komi            float64
	Trials          int
	factor          float64
	maxPlayoutDepth int
	minPlayout      int
}

func CreateEngine() Isana {
	return Isana{"", "", 0.0, 0, 0.31, 0, 1}
}

func (n *Isana) Ponder(pos *Position, s Stone) Move {

	defer Un(Trace("Isana#Ponder"))

	n.maxPlayoutDepth = int(Floor(float64(pos.Size.Capacity()) * 1.2))

	var wg WaitGroup
	for i := 0; i < n.Trials; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n.UCT(pos, s)
		}()
	}
	wg.Wait()

	selected, max := PassMove, -999
	for _, v := range pos.Moves {
		log.Printf("%v: %1.5f(%v)", v.String(), v.Rate, v.Games)
		if v.Games > max {
			selected = v
			max = v.Games
		}
	}

	log.Printf("selected... %v: %1.5f(%v)",
		selected.String(),
		selected.Rate,
		selected.Games)

	return selected
}

func (n *Isana) UCT(pos *Position, s Stone) float64 {
	maxUcb, selected := -999.0, pos.Size.Capacity()

	// If moves-slice is empty, create all moves.
	mutex.Lock()
	if len(pos.Moves) == 0 {
		for i := 0; i < pos.Size.Capacity(); i++ {
			mv := CreateMove(s, Vertex{i, pos.Size})
			pos.Moves = append(pos.Moves, *mv)
		}
		// Add pass
		pos.Moves = append(pos.Moves, PassMove)
	}
	mutex.Unlock()

	for i, v := range pos.Moves {
		if v != PassMove {
			_, ok := pos.PseudoMoveStrict(&v)
			if !ok {
				continue
			}
		}
		ucb := 0.0
		if v.Games == 0 {
			ucb = 10000
		} else {
			ucb = v.Rate + n.factor*Sqrt(Log10(float64(pos.Games))/float64(v.Games))
		}
		if ucb > maxUcb {
			maxUcb = ucb
			selected = i
		}
	}

	mv := &pos.Moves[selected]
	next, ok := pos.PseudoMoveStrict(mv)
  if !ok {
    next = pos
  }
	next.FixMove(mv)
	win := 0.0
	if mv.Games < n.minPlayout {
		win -= n.playout(CopyPosition(next), s.Opposite())
	} else {
		win -= n.UCT(next, s.Opposite())
	}

	mutex.Lock()
	mv.Rate = (mv.Rate*float64(mv.Games) + win) / float64(mv.Games+1)
	mv.Games++
	pos.Games++
	mutex.Unlock()

	return win
}

func (n *Isana) playout(pos Position, stone Stone) float64 {
	// Initialize probability dencities
	pos.CreateProbs()

	s, passed := stone, false
	depth := n.maxPlayoutDepth
	for depth > 0 {
		m := n.Inspiration(&pos, s)
		if m.Vertex == Outbound {
			if passed {
				break
			} else {
				passed = true
			}
		} else {
			passed = false
			pos.FixMove(m)
		}
		s = s.Opposite()
		depth--
	}
	return pos.Score(stone, n.Komi)
}

func (n *Isana) Inspiration(pos *Position, s Stone) *Move {
	if pos.TotalProbs <= 0 {
		return &PassMove
	}
	// loop at all candidates randomly
	i := pos.SearchProbIndex(Intn(pos.TotalProbs))
	// set the probability to 0.
	current := pos.ProbDencities[i] // keep the current value
	pos.UpdateProbs(i, 0)

	mv := CreateMove(s, Vertex{i, pos.Size})
	_, ok := pos.PseudoMoveStrict(mv)
	if ok {
		return mv
	}
	rec := n.Inspiration(pos, s)
	// if skip this move, revert the probability.
	pos.UpdateProbs(i, current)
	return rec
}
