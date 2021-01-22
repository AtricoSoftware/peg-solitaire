package board

import (
	"github.com/AtricoSoftware/peg-solitaire/api/direction"
	"github.com/AtricoSoftware/peg-solitaire/api/position"
	. "github.com/AtricoSoftware/peg-solitaire/xy"
)

func (b Board) GetMoves() MoveList {
	moves := make(MoveList, 0)
	b.forEachPegPosition(func(pos Position) {
		if b.holes[pos] {
			moves = append(moves, b.getMoves(pos)...)
		}
	})
	return moves
}

// Get moves for a single peg
func (b Board) getMoves(peg Position) MoveList {
	moves := make(MoveList, 0)
	direction.ForEachDirection(func(d direction.Direction) {
		pivot := position.Shift(peg, d)
		land := position.Shift(pivot, d)
		if present, valid := b.holes[pivot]; valid && present {
			if present, valid = b.holes[land]; valid && !present {
				move := NewMoveC(peg, d)
				moves = append(moves, move)
				// Follow on moves
				b2, _ := b.MakeMove(move)
				for _, next := range b2.getMoves(position.ShiftN(peg, d, 2)) {
					moves = append(moves, move.Join(next))
				}
			}
		}
	})
	return moves
}
