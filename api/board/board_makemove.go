package board

import (
	"errors"
	"fmt"
)

func (b Board) MakeMove(move Move) (Board, error) {
	b2 := b.Copy()
	current := move.Coord
	// Check source
	if !b.holes[current] {
		return b, errors.New(fmt.Sprintf("No peg to move at %v", move.Coord))
	}
	b2.holes[current] = false
	// Check pivot(s) and landings(s)
	for _, dir := range move.Directions {
		current = current.Shift(dir)
		if !b.holes[current] {
			return b, errors.New(fmt.Sprintf("No peg to jump at %v", current))
		}
		b2.holes[current] = false
		current = current.Shift(dir)
		if b.holes[current] {
			return b, errors.New(fmt.Sprintf("No landing space at %v", current))
		}
	}
	b2.holes[current] = true
	return b2, nil
}
