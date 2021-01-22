package unit_tests

import (
	"fmt"
	"testing"

	"github.com/atrico-go/core"
	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"

	"github.com/AtricoSoftware/peg-solitaire/api/board"
)

func Test_MakeMove_NoSource(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		Build()
	core.DisplayMultiline(b)
	move := board.NewMove(-1, -3, board.Down)

	// Act
	_, err := b.MakeMove(move)
	fmt.Printf("%v =>\n", move)

	// Assert
	Assert(t).That(err, is.NotNil, "Error")
	expected := fmt.Sprintf(`No peg to move at %v`, move.Position)
	Assert(t).That(err.Error(), is.EqualTo(expected), "Error message")
}

func Test_MakeMove_NoPivot(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		AddPeg(-1, -3).
		Build()
	core.DisplayMultiline(b)
	move := board.NewMove(-1, -3, board.Down)

	// Act
	_, err := b.MakeMove(move)
	fmt.Printf("%v =>\n", move)

	// Assert
	Assert(t).That(err, is.NotNil, "Error")
	expected := fmt.Sprintf(`No peg to jump at %v`, move.Position.Shift(board.Down))
	Assert(t).That(err.Error(), is.EqualTo(expected), "Error message")
}

func Test_MakeMove_NoLanding(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		AddPeg(-1, -3).
		AddPeg(-1, -2).
		AddPeg(-1, -1).
		Build()
	core.DisplayMultiline(b)
	move := board.NewMove(-1, -3, board.Down)

	// Act
	_, err := b.MakeMove(move)
	fmt.Printf("%v =>\n", move)

	// Assert
	Assert(t).That(err, is.NotNil, "Error")
	expected := fmt.Sprintf(`No landing space at %v`, move.Position.ShiftN(board.Down, 2))
	Assert(t).That(err.Error(), is.EqualTo(expected), "Error message")
}

func Test_MakeMove_NoNthPivot(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		AddPeg(-1, -3).
		AddPeg(-1, -2).
		Build()
	core.DisplayMultiline(b)
	move := board.NewMove(-1, -3, board.Down, board.Down)

	// Act
	_, err := b.MakeMove(move)
	fmt.Printf("%v =>\n", move)

	// Assert
	Assert(t).That(err, is.NotNil, "Error")
	expected := fmt.Sprintf(`No peg to jump at %v`, move.Position.ShiftN(board.Down, 3))
	Assert(t).That(err.Error(), is.EqualTo(expected), "Error message")
}

func Test_MakeMove_NoNthLanding(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		AddPeg(-1, -3).
		AddPeg(-1, -2).
		AddPeg(-1, 00).
		AddPeg(-1, +1).
		Build()
	core.DisplayMultiline(b)
	move := board.NewMove(-1, -3, board.Down, board.Down)

	// Act
	_, err := b.MakeMove(move)
	fmt.Printf("%v =>\n", move)

	// Assert
	Assert(t).That(err, is.NotNil, "Error")
	expected := fmt.Sprintf(`No landing space at %v`, move.Position.ShiftN(board.Down, 4))
	Assert(t).That(err.Error(), is.EqualTo(expected), "Error message")
}

func Test_MakeMove_Single(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		AddPeg(-1, -3).
		AddPeg(-1, -2).
		Build()
	core.DisplayMultiline(b)
	move := board.NewMove(-1, -3, board.Down)

	// Act
	b2, err := b.MakeMove(move)
	fmt.Printf("%v =>\n", move)
	core.DisplayMultiline(b2)

	// Assert
	Assert(t).That(err, is.Nil, "No Error")
	expected := board.NewBoardBuilder().
		AddPeg(-1, -1).
		Build()
	Assert(t).That(b2.Id(), is.EqualTo(expected.Id()), "Correct board")
}

func Test_MakeMove_Multiple(t *testing.T) {
	// Arrange
	b := board.NewBoardBuilder().
		AddPeg(-1, -3).
		AddPeg(-1, -2).
		AddPeg(-1, 00).
		AddPeg(-1, +2).
		AddPeg(00, +3).
		Build()
	core.DisplayMultiline(b)
	move := board.NewMove(-1, -3, board.Down, board.Down, board.Down, board.Right)

	// Act
	b2, err := b.MakeMove(move)
	fmt.Printf("%v =>\n", move)
	core.DisplayMultiline(b2)

	// Assert
	Assert(t).That(err, is.Nil, "No Error")
	expected := board.NewBoardBuilder().
		AddPeg(+1, +3).
		Build()
	Assert(t).That(b2.Id(), is.EqualTo(expected.Id()), "Correct board")
}
