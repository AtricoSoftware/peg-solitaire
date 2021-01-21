package board

import (
	"strings"
)

//   0 1 2 3 4 5 6
// 0     @ @ @
// 1     @ @ @
// 2 @ @ @ @ @ @ @
// 3 @ @ @ o @ @ @
// 4 @ @ @ @ @ @ @
// 5     @ @ @
// 6     @ @ @

type Board struct {
	holes map[Coord]bool
}

func NewBoard() Board {
	return NewBoardFromId("111111111111111101111111111111111")
}

func NewBoardFromId(id string) Board {
	b := Board{holes: make(map[Coord]bool, 33)}
	idx := 0
	forEachPegCoord(func(pos Coord) {
		b.holes[pos] = id[idx] == '1'
		idx++
	})
	return b
}

func (b Board) Id() string {
	id := strings.Builder{}
	forEachPegCoord(func(pos Coord) {
		if b.holes[pos] {
			id.WriteRune('1')
		} else {
			id.WriteRune('0')
		}
	})
	return id.String()
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------

func (b Board) Copy() Board {
	return NewBoardFromId(b.Id())
}
