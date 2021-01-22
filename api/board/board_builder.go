package board

import "fmt"

type Builder interface {
	AddPeg(x, y int) Builder
	Build() Board
}

func NewBoardBuilder() Builder {
	b := NewBoardFromId("000000000000000000000000000000000")
	return &b
}

func (b *Board) AddPeg(x, y int) Builder {
	pos := Position{x, y}
	if v,ok := b.holes[pos]; !ok {
		panic(fmt.Sprintf(`Invalid coordinate: "%v"`, pos))
	} else if v {
		panic(fmt.Sprintf(`Peg already present: "%v"`, pos))
	}
	b.holes[pos] = true
	return b
}

func (b *Board) Build() Board {
	return *b
}
