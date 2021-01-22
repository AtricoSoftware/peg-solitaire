package board

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
	forEachDirection(func(d Direction) {
		pivot := peg.Shift(d)
		land := pivot.Shift(d)
		if present, valid := b.holes[pivot]; valid && present {
			if present, valid = b.holes[land]; valid && !present {
				move := NewMoveC(peg, d)
				moves = append(moves, move)
				// Follow on moves
				b2, _ := b.MakeMove(move)
				for _, next := range b2.getMoves(peg.ShiftN(d, 2)) {
					moves = append(moves, move.Join(next))
				}
			}
		}
	})
	return moves
}
