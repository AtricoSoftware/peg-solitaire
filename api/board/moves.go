package board

import "fmt"

type Move struct {
	Coord
	Directions []Direction
}

type MoveList []Move

func NewMoveC(pos Coord, dirs ...Direction) Move {
	return Move{pos, dirs}
}

func NewMove(x, y int, dirs ...Direction) Move {
	return NewMoveC(Coord{x, y}, dirs...)
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------
func (m Move) String() string {
	return fmt.Sprintf("%v:%v", m.Coord, m.Directions)
}

func (m Move) Join(move Move) Move {
	return Move{m.Coord, append(m.Directions, move.Directions...)}
}
