package board

import (
	. "github.com/AtricoSoftware/peg-solitaire/xy"
)

type Board struct {
	holes map[Position]bool
}

type BoardId uint64

func NewStandardBoard() Board {
	return NewBoardFromId(0x1FFFEFFFF)
}

func NewEmptyBoard() Board {
	b := Board{holes: make(map[Position]bool, 33)}
	for y := -3; y <= +3; y++ {
		for x := -3; x <= +3; x++ {
			if (intAbs(x) <= 3 && intAbs(y) <= 1) || (intAbs(x) <= 1 && intAbs(y) <= 3) {
				b.holes[Position{X: x, Y: y}] = false
			}
		}
	}
	return b
}

func NewBoardFromId(id BoardId) Board {
	b := NewEmptyBoard()
	b.forEachPegPositionReverse(func(pos Position) {
		b.holes[pos] = id&1 == 1
		id >>= 1

	})
	return b
}

func (b Board) Id() BoardId {
	id := BoardId(0)
	b.forEachPegPosition(func(pos Position) {
		id <<= 1
		if b.holes[pos] {
			id |= 1
		}
	})
	return id
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
