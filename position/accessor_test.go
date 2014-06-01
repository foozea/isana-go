package position

import (
	"testing"

	. "github.com/foozea/isana/board/size"
	. "github.com/foozea/isana/board/stone"
	. "github.com/foozea/isana/board/vertex"
)

func init() {
	pos = CreatePosition(B9x9)
}

func TestSetAndGetStone(t *testing.T) {
	var actual, expected Stone
	const msg string = "Set,GetStone / failed to set/get a stone. expected : %v, but %v"
	pos := CreatePosition(B9x9)
	vx1 := StringToVertex("E3", B9x9)
	pos.SetStone(Black, vx1)
	actual = pos.GetStone(vx1)
	expected = Black
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	vx2 := StringToVertex("E4", B9x9)
	pos.SetStone(White, vx2)
	actual = pos.GetStone(vx2)
	expected = White
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = pos.GetStone(vx1)
	expected = Black
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
	actual = pos.GetStone(Outbound)
	expected = Wall
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}

func TestEmpties(t *testing.T) {
	var actual, expected int
	const msg string = "Empties / failed to get empties. expected : %v, but %v"
	actual = len(pos.Empties())
	expected = 81
	if actual != expected {
		t.Errorf(msg, expected, actual)
	}
}
