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

func TestGetAndCreateString(t *testing.T) {
	const msg string = "Create,GetString / failed to set/get string. expected : %v, but %v"
	vx := StringToVertex("E5", pos.Size)
	pos.SetStone(Black, vx)
	pos.CreateString(Black, vx)
	id, g := pos.GetString(vx)
	if id != 1 {
		t.Errorf(msg, 1, id)
	}
	if g == nil {
		t.Errorf(msg, g, nil)
	}
	if g != nil && g.Stone != Black {
		t.Errorf(msg, g.Stone, Black)
	}
}

func TestUpdateString(t *testing.T) {
	const msg string = "UpdateString / failed to update strings. expected : %v, but %v"
	vx1 := StringToVertex("E5", pos.Size)
	vx2 := StringToVertex("E6", pos.Size)
	vx3 := StringToVertex("E7", pos.Size)
	pos.SetStone(Black, vx1)
	pos.SetStone(Black, vx2)
	pos.SetStone(White, vx3)
	pos.UpdateStrings()
	id1, g1 := pos.GetString(vx1)
	id2, g2 := pos.GetString(vx2)
	id3, g3 := pos.GetString(vx3)

	if id1 != id2 || *g1 != *g2 {
		t.Errorf(msg, *g1, *g2)
	}
	if id3 == id1 || *g3 == *g1 {
		t.Errorf(msg, *g3, *g1)
	}
	if g1.Stone != Black {
		t.Errorf(msg, g3.Stone, Black)
	}
	if g3.Stone != White {
		t.Errorf(msg, g3.Stone, White)
	}

	pos.GoStringDump()
}
