package board

import "fmt"

type Coord struct {
	X, Y int
}

func (c Coord) String() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}

func (c Coord) Shift(d Direction) Coord {
	return c.ShiftN(d, 1)
}

func (c Coord) ShiftN(d Direction, n int) Coord {
	switch d {
	case Up:
		return Coord{c.X, c.Y - n}
	case Down:
		return Coord{c.X, c.Y + n}
	case Left:
		return Coord{c.X - n, c.Y}
	case Right:
		return Coord{c.X + n, c.Y}
	}
	panic(fmt.Sprintf(`Invalid direction: "%v"`, d))

}

func forEachPegCoord(f func(Coord)) {
	for y := 0; y < 7; y++ {
		for x := 0; x < 7; x++ {
			c := Coord{x, y}
			if c.isValid() {
				f(c)
			}
		}
	}
}

func (c Coord) isValid() bool {
	return (2 <= c.X && c.X <= 4 && 0 <= c.Y && c.Y <= 6) || (0 <= c.X && c.X <= 6 && 2 <= c.Y && c.Y <= 4)
}
