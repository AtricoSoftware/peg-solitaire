package board

import "fmt"

type Position struct {
	X, Y int
}

func (c Position) String() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}

func (c Position) Shift(d Direction) Position {
	return c.ShiftN(d, 1)
}

func (c Position) ShiftN(d Direction, n int) Position {
	switch d {
	case Up:
		return Position{c.X, c.Y - n}
	case Down:
		return Position{c.X, c.Y + n}
	case Left:
		return Position{c.X - n, c.Y}
	case Right:
		return Position{c.X + n, c.Y}
	}
	panic(fmt.Sprintf(`Invalid direction: "%v"`, d))
}

func (c Position) Less(rhs Position) bool {
	if c.Y < rhs.Y {
		return true
	}
	if c.Y > rhs.Y {
		return false
	}
	return c.X < rhs.X
}