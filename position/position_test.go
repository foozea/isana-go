package position

import (
	"testing"

	. "github.com/foozea/isana/board/size"
	. "github.com/foozea/isana/board/stone"
	. "github.com/foozea/isana/board/vertex"
	. "github.com/foozea/isana/position/move"
)

var (
	pos Position
)

func init() {
	pos = CreatePosition(B9x9)
}

func TestCountLiberty(t *testing.T) {
	var actual, expected int
	const msg string = "CountStringLiberty / counted number is not correct. expected : %v, but %v"
	pos := CreatePosition(B9x9)
	v1 := StringToVertex("A8", B9x9)
	pos.FixMove(CreateMove(Black, v1))
	v3 := StringToVertex("A9", B9x9)
	pos.FixMove(CreateMove(White, v3))
	actual = pos.CountStringLiberty(pos.GetString(v3))
	expected = 1
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	v2 := StringToVertex("A9", B9x9)
	pos.FixMove(CreateMove(Black, v2))
	actual = pos.CountStringLiberty(pos.GetString(v2))
	expected = 1
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}

	pos = CreatePosition(B9x9)
	v1 = StringToVertex("J1", B9x9)
	v2 = StringToVertex("J2", B9x9)
	v3 = StringToVertex("J3", B9x9)
	v4 := StringToVertex("J4", B9x9)
	v5 := StringToVertex("H4", B9x9)
	v6 := StringToVertex("H1", B9x9)
	v7 := StringToVertex("H2", B9x9)
	v8 := StringToVertex("H3", B9x9)
	v9 := StringToVertex("G4", B9x9)
	v10 := StringToVertex("H5", B9x9)
	v11 := StringToVertex("J5", B9x9)
	pos.FixMove(CreateMove(Black, v1))
	pos.FixMove(CreateMove(Black, v2))
	pos.FixMove(CreateMove(Black, v3))
	pos.FixMove(CreateMove(Black, v4))
	pos.FixMove(CreateMove(Black, v5))
	pos.FixMove(CreateMove(White, v6))
	pos.FixMove(CreateMove(White, v7))
	pos.FixMove(CreateMove(White, v8))
	pos.FixMove(CreateMove(White, v9))
	pos.FixMove(CreateMove(White, v10))
	actual = pos.CountStringLiberty(pos.GetString(v5))
	expected = 1
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	pos.FixMove(CreateMove(White, v11))

	pos = CreatePosition(B9x9)
	v1 = StringToVertex("A5", B9x9)
	v2 = StringToVertex("A8", B9x9)
	v3 = StringToVertex("B6", B9x9)
	v4 = StringToVertex("B7", B9x9)
	v5 = StringToVertex("A6", B9x9)
	v6 = StringToVertex("A7", B9x9)
	v7 = StringToVertex("B5", B9x9)
	v8 = StringToVertex("B4", B9x9)
	v9 = StringToVertex("C5", B9x9)
	pos.FixMove(CreateMove(White, v1))
	pos.FixMove(CreateMove(White, v2))
	pos.FixMove(CreateMove(Black, v5))
	pos.FixMove(CreateMove(Black, v6))
	pos.FixMove(CreateMove(White, v8))
	pos.FixMove(CreateMove(Black, v7))
	pos.FixMove(CreateMove(White, v4))
	pos.FixMove(CreateMove(White, v9))
	pos.FixMove(CreateMove(White, v3))

	pos = CreatePosition(B9x9)
	v1 = StringToVertex("A8", B9x9)
	v2 = StringToVertex("B8", B9x9)
	v3 = StringToVertex("C9", B9x9)
	v4 = StringToVertex("A9", B9x9)
	v5 = StringToVertex("B9", B9x9)
	pos.FixMove(CreateMove(Black, v2))
	pos.FixMove(CreateMove(Black, v3))
	pos.FixMove(CreateMove(White, v4))
	pos.FixMove(CreateMove(White, v5))
	pos.FixMove(CreateMove(Black, v1))
}

func TestScore(t *testing.T) {
	var actual, expected float64
	const msg string = "Score / failed to count valid score. expected : %v, but %v"
	next, ok := pos.PseudoMove(CreateMove(Black, StringToVertex("E4", B9x9)), false)
	if !ok {
		t.Errorf(msg, true, ok)
	}
	next.FixMove(CreateMove(White, StringToVertex("E3", B9x9)))
	next.FixMove(CreateMove(Black, StringToVertex("E2", B9x9)))
	next.FixMove(CreateMove(Black, StringToVertex("D3", B9x9)))
	next.FixMove(CreateMove(Black, StringToVertex("F3", B9x9)))
	next.FixMove(CreateMove(White, StringToVertex("A3", B9x9)))
	next.FixMove(CreateMove(White, StringToVertex("A4", B9x9)))

	actual = next.Score(Black, 0.0)
	expected = 1
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}

	pos.Dump()
}
