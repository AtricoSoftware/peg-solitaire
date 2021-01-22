package board

import (
	"fmt"

	"github.com/AtricoSoftware/peg-solitaire/api/direction"
	. "github.com/AtricoSoftware/peg-solitaire/xy"
)

type Move struct {
	Position
	Directions []direction.Direction
}

type MoveList []Move

func NewMoveC(pos Position, dirs ...direction.Direction) Move {
	return Move{pos, dirs}
}

func NewMove(x, y int, dirs ...direction.Direction) Move {
	return NewMoveC(Position{X: x, Y: y}, dirs...)
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
