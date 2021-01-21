package board

import "fmt"

type Direction string

const (
	Up    Direction = "Up"
	Down  Direction = "Down"
	Left  Direction = "Left"
	Right Direction = "Right"
)

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------
func (d Direction) Opposite() Direction {
	switch d {
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	}
	panic(fmt.Sprintf(`Invalid direction: "%v"`, d))
}

func forEachDirection(f func(Direction)) {
	for _, d := range []Direction{Up, Down, Left, Right} {
		f(d)
	}
}
