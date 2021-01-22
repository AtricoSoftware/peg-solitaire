package board

import "fmt"

type Move struct {
	Position
	Directions []Direction
}

type MoveList []Move

func NewMoveC(pos Position, dirs ...Direction) Move {
	return Move{pos, dirs}
}

func NewMove(x, y int, dirs ...Direction) Move {
	return NewMoveC(Position{x, y}, dirs...)
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------
func (m Move) String() string {
	return fmt.Sprintf("%v:%v", m.Position, m.Directions)
}

func (m Move) Join(move Move) Move {
	return Move{m.Position, append(m.Directions, move.Directions...)}
}
