package board

import (
	"errors"
	"fmt"

	"github.com/AtricoSoftware/peg-solitaire/api/position"
)

func (b Board) MakeMove(move Move) (Board, error) {
	b2 := b.Copy()
	current := move.Position
	// Check source
	if !b2.holes[current] {
		return b, errors.New(fmt.Sprintf("No peg to move at %v", move.Position))
	}
	b2.holes[current] = false
	// Check pivot(s) and landings(s)
	for _, dir := range move.Directions {
		current = position.Shift(current, dir)
		if !b2.holes[current] {
			return b, errors.New(fmt.Sprintf("No peg to jump at %v", current))
		}
		b2.holes[current] = false
		current = position.Shift(current, dir)
		if b2.holes[current] {
			return b, errors.New(fmt.Sprintf("No landing space at %v", current))
		}
	}
	b2.holes[current] = true
	return b2, nil
}
