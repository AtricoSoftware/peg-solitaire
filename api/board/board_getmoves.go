package board

func (b Board) GetMoves() []Move {
	moves := make([]Move, 0)
	forEachPegCoord(func(pos Coord) {
		if b.holes[pos] {
			moves = append(moves, b.getMoves(pos)...)
		}
	})
	return moves
}


// Get moves for a single peg
func (b Board) getMoves(peg Coord) []Move {
	moves := make([]Move, 0)
	forEachDirection(func(d Direction) {
		pivot := peg.Shift(d)
		land := pivot.Shift(d)
		if pivot.isValid() && b.holes[pivot] && land.isValid() && !b.holes[land] {
			move := NewMoveC(peg, d)
			moves = append(moves, move)
			// Follow on moves
			b2,_ := b.MakeMove(move)
			for _,next := range b2.getMoves(peg.ShiftN(d,2)) {
				moves = append(moves, move.Join(next))
			}
		}
	})
	return moves
}
