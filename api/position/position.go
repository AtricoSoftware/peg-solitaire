package position

import (
	"github.com/AtricoSoftware/peg-solitaire/api/direction"
	. "github.com/AtricoSoftware/peg-solitaire/xy"
)

func Shift(p Position, d direction.Direction) Position {
	return ShiftN(p,d, 1)
}

func ShiftN(p Position, d direction.Direction, n int) Position {
	return p.Add(Vector(d).Mult(n))
}

func Less(lhs Position, rhs Position) bool {
	if lhs.Y < rhs.Y {
		return true
	}
	if lhs.Y > rhs.Y {
		return false
	}
	return lhs.X < rhs.X
}