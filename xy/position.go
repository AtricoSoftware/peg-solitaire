package xy

type Position XyPair

func (lhs Position) String() string {
	return XyPair(lhs).StringRound()
}

func (lhs Position) Add(rhs Vector) Position {
	return Position(XyPair(lhs).Add(XyPair(rhs)))
}

func (lhs Position) Sub(rhs Vector) Position {
	return Position(XyPair(lhs).Sub(XyPair(rhs)))
}
