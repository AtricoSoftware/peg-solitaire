package board

import (
	"strings"
)

type Board struct {
	holes map[Position]bool
}

func NewBoard() Board {
	return NewBoardFromId("111111111111111101111111111111111")
}

func NewEmptyBoard() Board {
	b := Board{holes: make(map[Position]bool, 33)}
	for y := -3; y <= +3; y++ {
		for x := -3; x <= +3; x++ {
			if (intAbs(x) <= 3 && intAbs(y) <= 1) || (intAbs(x) <= 1 && intAbs(y) <= 3) {
				b.holes[Position{x, y}] = false
			}
		}
	}
	return b
}

func NewBoardFromId(id string) Board {
	b := NewEmptyBoard()
	idx := 0
	b.forEachPegPosition(func(pos Position) {
		b.holes[pos] = id[idx] == '1'
		idx++
	})
	return b
}

func (b Board) Id() string {
	id := strings.Builder{}
	b.forEachPegPosition(func(pos Position) {
		if b.holes[pos] {
			id.WriteRune('1')
		} else {
			id.WriteRune('0')
		}
	})
	return id.String()
}

func (b Board) IsSolved() bool {
	return b.PegsRemaining() == 1
}

func (b Board) PegsRemaining() int {
	total := 0
	b.forEachPegPosition(func(pos Position) {
		if b.holes[pos] {
			total++
		}
	})
	return total
}

func (b Board) Copy() Board {
	return NewBoardFromId(b.Id())
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------

func intAbs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
